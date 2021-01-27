package manage

import (
	"context"
	"github.com/fyuan1316/operatorlib/api"
	"github.com/fyuan1316/operatorlib/event"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/shell"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	//"github.com/fyuan1316/asm-operator/pkg/logging"
	pkgerrors "github.com/pkg/errors"
	//"github.com/sirupsen/logrus"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var (
	//logger = logging.RegisterScope("controller.oprlib")
	DefaultBackoff = wait.Backoff{Duration: time.Millisecond * 10, Factor: 2, Steps: 3}
)

func (m *OperatorManage) Reconcile(instance model.CommonOperator, provisionStages, deletionStages [][]model.ExecuteItem) (ctrl.Result, error) {
	var (
		oCtx *model.OperatorContext
		err  error
	)
	if oCtx, err = model.NewOperatorContext(m.K8sClient, m.Recorder, instance); err != nil {
		m.Recorder.Event(instance, event.WarningEvent, ParseParamsError, err.Error())
		return ctrl.Result{}, pkgerrors.WithStack(err)
	}

	if m.Options.FinalizerID != "" && instance.GetDeletionTimestamp().IsZero() {
		//logger.Debug("sync cr")
		if !ContainsString(instance.GetFinalizers(), m.Options.FinalizerID) {
			finalizers := append(instance.GetFinalizers(), m.Options.FinalizerID)
			instance.SetFinalizers(finalizers)
			if err = m.K8sClient.Update(context.Background(), instance.DeepCopyObject()); err != nil {
				return ctrl.Result{}, pkgerrors.Wrap(err, "could not add finalizer")
			}
			return ctrl.Result{
				RequeueAfter: time.Second * 1,
			}, nil
		}
	} else if !instance.GetDeletionTimestamp().IsZero() {
		//logger.Debug("delete cr")
		if len(deletionStages) > 0 {
			//标识当前为删除过程，相应信息记录在status.deleteConditions
			oCtx.DoDeletion()
			if sc := m.ProcessStages(deletionStages, oCtx); sc.Err() != nil {
				// update partial status
				if updErr := m.Options.StatusUpdater(oCtx, sc); updErr != nil {
					return ctrl.Result{}, updErr
				}
				return ctrl.Result{}, sc.Err()
			}
		}
		//deletion success
		f := RemoveString(instance.GetFinalizers(), m.Options.FinalizerID)
		instance.SetFinalizers(f)
		if err = m.K8sClient.Update(context.Background(), instance.DeepCopyObject()); err != nil {
			return reconcile.Result{}, pkgerrors.Wrap(err, "could not remove finalizer")
		}
		return ctrl.Result{}, nil
	}
	if len(provisionStages) == 0 {
		return ctrl.Result{}, nil
	}
	//sync
	//标识当前为安装过程，相应信息记录在status.installConditions
	oCtx.DoProvision()
	syncSc := m.ProcessStages(provisionStages, oCtx)
	if syncSc.Err() != nil {
		// update partial status
		if updErr := m.Options.StatusUpdater(oCtx, syncSc); updErr != nil {
			return ctrl.Result{}, updErr
		}
		return ctrl.Result{}, syncSc.Err()
	}

	isReady, isHealthy, err := m.DoHealthCheck(provisionStages, oCtx)
	if err != nil {
		return ctrl.Result{}, err
	}
	// append healthy check state
	syncSc.OperatorStatus.SetState(isReady, isHealthy)
	// update total status
	if err = m.Options.StatusUpdater(oCtx, syncSc); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (m *OperatorManage) DoHealthCheck(stages [][]model.ExecuteItem, oCtx *model.OperatorContext) (bool, bool, error) {
	var readyCheckNum, readyNum, healthyCheckNum, healthyNum int
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := model.CanDoHealthCheck(item); ok {
				readyCheckNum += 1
				if ref.IsReady(oCtx) {
					readyNum += 1
				}
				healthyCheckNum += 1
				if ref.IsHealthy(oCtx) {
					healthyNum += 1
				}
			}
		}
	}
	return readyCheckNum == readyNum, healthyCheckNum == healthyNum, nil
}

func (m *OperatorManage) ProcessStages(stages [][]model.ExecuteItem, oCtx *model.OperatorContext) *model.StatusContext {
	sc := model.NewStatusContext(*oCtx)
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := model.CanDoPreCheck(item); ok {
				err := loopUntil(context.Background(), defualtLoopSetting.Interval*time.Second, defualtLoopSetting.MaxRetries, ref.PreCheck, oCtx)
				if pkgerrors.Is(err, shell.InternalIgnoreShellScriptError) {
					continue
				}
				cond := sc.StageCondition(oCtx.OperationType(), api.OperationStages.PreCheck, item.Name())
				cond.SetLastTransitionTime(metav1.Now())
				if err != nil {
					cond.SetFailed().
						SetReason("PreCheck failed").
						SetMessage(err.Error())
					sc.SetError(err)
					return sc
				}

			}
		}
		for _, item := range items {
			if ref, ok := model.CanDoPreRun(item); ok {
				//logger.Debugf("run prerun")
				err := ref.PreRun(oCtx)
				if pkgerrors.Is(err, shell.InternalIgnoreShellScriptError) {
					continue
				}
				cond := sc.StageCondition(oCtx.OperationType(), api.OperationStages.PreRun, item.Name())
				cond.SetLastTransitionTime(metav1.Now())
				if err != nil {
					cond.SetFailed().
						SetReason("PreRun failed").
						SetMessage(err.Error())
					sc.SetError(err)
					return sc
				}
			}
		}
		for _, item := range items {
			err := item.Run(oCtx)
			if pkgerrors.Is(err, shell.InternalIgnoreShellScriptError) {
				continue
			}
			cond := sc.StageCondition(oCtx.OperationType(), api.OperationStages.Run, item.Name())
			cond.SetLastTransitionTime(metav1.Now())
			if err != nil {
				cond.SetFailed().
					SetReason("Run failed").
					SetMessage(err.Error())
				sc.SetError(err)
				return sc
			}
		}

		for _, item := range items {
			if ref, ok := model.CanDoPostRun(item); ok {
				//logger.Debugf("run postrun")
				err := ref.PostRun(oCtx)
				if pkgerrors.Is(err, shell.InternalIgnoreShellScriptError) {
					continue
				}
				cond := sc.StageCondition(oCtx.OperationType(), api.OperationStages.PostRun, item.Name())
				cond.SetLastTransitionTime(metav1.Now())
				if err != nil {
					cond.SetFailed().
						SetReason("PostRun failed").
						SetMessage(err.Error())
					sc.SetError(err)
					return sc
				}
			}
		}
		for _, item := range items {
			if ref, ok := model.CanDoPostCheck(item); ok {
				//logger.Debugf("run postcheck")
				err := loopUntil(context.Background(), defualtLoopSetting.Interval*time.Second, defualtLoopSetting.MaxRetries, ref.PostCheck, oCtx)
				if pkgerrors.Is(err, shell.InternalIgnoreShellScriptError) {
					continue
				}
				cond := sc.StageCondition(oCtx.OperationType(), api.OperationStages.PostCheck, item.Name())
				cond.SetLastTransitionTime(metav1.Now())
				if err != nil {
					cond.SetFailed().
						SetReason("PostCheck failed").
						SetMessage(err.Error())
					sc.SetError(err)
					return sc
				}
			}
		}
	}
	return sc
}

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
