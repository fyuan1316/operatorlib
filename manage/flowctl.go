package manage

import (
	"context"
	"github.com/fyuan1316/operatorlib/manage/model"

	//"github.com/fyuan1316/asm-operator/pkg/logging"
	pkgerrors "github.com/pkg/errors"
	//"github.com/sirupsen/logrus"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var (
//logger = logging.RegisterScope("controller.oprlib")
)

func (m *OperatorManage) Reconcile(instance model.CommonOperator, provisionStages, deletionStages [][]model.ExecuteItem) (ctrl.Result, error) {
	//logger.SetOutputLevel(logrus.DebugLevel)
	var (
		params map[string]interface{}
		err    error
	)
	if params, err = instance.GetOperatorParams(); err != nil {
		return ctrl.Result{}, pkgerrors.Wrap(err, "parse spec params error")
	}
	oCtx := model.OperatorContext{
		K8sClient:      m.K8sClient,
		Recorder:       m.Recorder,
		Instance:       instance,
		OperatorParams: params,
	}

	if m.Options.FinalizerID != "" && instance.GetDeletionTimestamp().IsZero() {
		//logger.Debug("sync cr")
		if !ContainsString(instance.GetFinalizers(), m.Options.FinalizerID) {
			finalizers := append(instance.GetFinalizers(), m.Options.FinalizerID)
			instance.SetFinalizers(finalizers)
			if err := m.K8sClient.Update(context.Background(), instance.DeepCopyObject()); err != nil {
				return ctrl.Result{}, pkgerrors.Wrap(err, "could not add finalizer to config")
			}
			return ctrl.Result{
				RequeueAfter: time.Second * 1,
			}, nil
		}
	} else if !instance.GetDeletionTimestamp().IsZero() {
		//logger.Debug("delete cr")
		if len(deletionStages) > 0 {
			if err := m.ProcessStages(deletionStages, oCtx); err != nil {
				return ctrl.Result{}, err
			}
		}
		f := RemoveString(instance.GetFinalizers(), m.Options.FinalizerID)
		instance.SetFinalizers(f)
		//logger.Debugf("cr Finalizers=%v", instance.GetFinalizers())
		if err := m.K8sClient.Update(context.Background(), instance.DeepCopyObject()); err != nil {
			return reconcile.Result{}, pkgerrors.Wrap(err, "could not remove finalizer from config")
		}
		return ctrl.Result{}, nil
	}
	if len(provisionStages) == 0 {
		return ctrl.Result{}, nil
	}
	//sync
	if err := m.ProcessStages(provisionStages, oCtx); err != nil {
		return ctrl.Result{}, err
	}
	if err := m.DoHealthCheck(provisionStages, oCtx); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil

}

func (m *OperatorManage) DoHealthCheck(stages [][]model.ExecuteItem, oCtx model.OperatorContext) error {
	//oCtx := OperatorContext{
	//	K8sClient:      m.K8sClient,
	//	Recorder:       m.Recorder,
	//	OperatorParams: instance.GetOperatorParams(),
	//}
	var readyCheckNum, readyNum, healthyCheckNum, healthyNum int
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := model.CanDoHealthCheck(item); ok {
				//logger.Debugf("run HealthCheck")
				readyCheckNum += 1
				if ref.IsReady(&oCtx) {
					readyNum += 1
				}
				healthyCheckNum += 1
				if ref.IsHealthy(&oCtx) {
					healthyNum += 1
				}
			}
		}
	}
	// if some task needs report its states, we update operator cr's status.state
	if readyCheckNum > 0 {
		if err := m.Options.StatusUpdater(oCtx.Instance.DeepCopyObject(), m.K8sClient)(readyCheckNum == readyNum, healthyCheckNum == healthyNum); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (m *OperatorManage) ProcessStages(stages [][]model.ExecuteItem, oCtx model.OperatorContext) error {

	for _, items := range stages {
		for _, item := range items {
			if ref, ok := model.CanDoPreCheck(item); ok {
				//logger.Debugf("run precheck")
				if err := loopUntil(context.Background(), 5*time.Second, 3, ref.PreCheck, &oCtx); err != nil {
					return err
				}
			}
		}
		for _, item := range items {
			if ref, ok := model.CanDoPreRun(item); ok {
				//logger.Debugf("run prerun")
				if err := ref.PreRun(&oCtx); err != nil {
					return err
				}
			}
		}
		for _, item := range items {
			if err := item.Run(&oCtx); err != nil {
				return err
			}
		}

		for _, item := range items {
			if ref, ok := model.CanDoPostRun(item); ok {
				//logger.Debugf("run postrun")
				if err := ref.PostRun(&oCtx); err != nil {
					return err
				}
			}
		}
		for _, item := range items {
			if ref, ok := model.CanDoPostCheck(item); ok {
				//logger.Debugf("run postcheck")
				if err := loopUntil(context.Background(), 5*time.Second, 3, ref.PostCheck, &oCtx); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
