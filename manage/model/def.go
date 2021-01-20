package model

import (
	"github.com/fyuan1316/operatorlib/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Object interface {
	runtime.Object
	metav1.Object
}

type CommonOperator interface {
	runtime.Object
	metav1.Object
	GetOperatorParams() (map[string]interface{}, error)
}
type OperatorContext struct {
	K8sClient      client.Client
	Recorder       record.EventRecorder
	Instance       CommonOperator
	operationType  OperationType
	OperatorParams map[string]interface{}
}

func (oc *OperatorContext) DoProvision() {
	oc.operationType = Operations.Provision
}
func (oc *OperatorContext) DoDeletion() {
	oc.operationType = Operations.Deletion
}
func (oc OperatorContext) OperationType() OperationType {
	return oc.operationType
}

type StatusContext struct {
	err error
	api.OperatorStatus
	operationType OperationType
}

func NewStatusContext(oCtx OperatorContext) *StatusContext {
	return &StatusContext{operationType: oCtx.operationType}
}
func (sc *StatusContext) Err() error {
	return sc.err
}

func (sc *StatusContext) OperationType() OperationType {
	return sc.operationType
}

func (sc *StatusContext) SetError(err error) *StatusContext {
	sc.err = err
	return sc
}

func (sc *StatusContext) SetOperationType(t OperationType) *StatusContext {
	sc.operationType = t
	return sc
}

func findConditionByName(conditions []*api.Condition, name string) *api.Condition {
	if conditions == nil {
		return nil
	}
	for i := range conditions {
		cond := conditions[i]
		if name == cond.Name {
			return cond
		}
	}
	return nil
}
func (sc *StatusContext) StageCondition(operation OperationType, stage api.OperationStage, name string) *api.Condition {
	var stageConditions *map[api.OperationStage]api.TaskCondition
	switch operation {
	case Operations.Deletion:
		if sc.DeleteConditions == nil {
			sc.DeleteConditions = make(map[api.OperationStage]api.TaskCondition, 0)
		}
		stageConditions = &sc.DeleteConditions
	default:
		if sc.InstallConditions == nil {
			sc.InstallConditions = make(map[api.OperationStage]api.TaskCondition, 0)
		}
		stageConditions = &sc.InstallConditions
	}

	conditions, exists := (*stageConditions)[stage]
	if !exists {
		c := NewDefaultCondition(name)
		conditions = append(conditions, c)
		(*stageConditions)[stage] = conditions
		return c
	}
	if cond := findConditionByName(conditions, name); cond != nil {
		return cond
	}
	c := NewDefaultCondition(name)
	conditions = append(conditions, c)
	(*stageConditions)[stage] = conditions
	return c
}

func NewDefaultCondition(name string) *api.Condition {
	c := &api.Condition{}
	c.Name = name
	c.Status = api.ConditionState.Succeed
	return c
}
