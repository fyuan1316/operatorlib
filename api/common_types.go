package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// +k8s:deepcopy-gen=false

type OperatorSpec struct {
	Namespace  string `json:"namespace,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

func (in OperatorSpec) GetOperatorParams() (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := yaml.Unmarshal([]byte(in.Parameters), &m); err != nil {
		return nil, err
	}

	return m, nil
}
func (in OperatorSpec) GetInstalledNamespace() string {
	return in.Namespace
}

type OperatorState string

var OperatorStates = struct {
	NotReady OperatorState
	Ready    OperatorState
	Health   OperatorState
}{"NotReady", "Ready", "Health"}

type TaskCondition []*Condition

// +k8s:openapi-gen=true
type OperatorStatus struct {
	State             OperatorState                    `json:"state,omitempty"`
	InstallConditions map[OperationStage]TaskCondition `json:"installConditions,omitempty"`
	DeleteConditions  map[OperationStage]TaskCondition `json:"deleteConditions,omitempty"`
}

func (in *OperatorStatus) DeepCopyInto(out *OperatorStatus) {
	if in == nil {
		return
	}
	out.State = in.State
	if len(in.InstallConditions) > 0 {
		m := make(map[OperationStage]TaskCondition, len(in.InstallConditions))
		for k, v := range in.InstallConditions {
			//length := len(v)
			slice := make(TaskCondition, 0)
			for _, vv := range v {
				slice = append(slice, vv)
			}
			m[k] = slice
		}
		out.InstallConditions = m
	}
	if len(in.DeleteConditions) > 0 {
		m := make(map[OperationStage]TaskCondition, len(in.DeleteConditions))
		for k, v := range in.DeleteConditions {
			length := len(v)
			slice := make(TaskCondition, length)
			for _, vv := range v {
				slice = append(slice, vv)
			}
			m[k] = slice
		}
		out.DeleteConditions = m
	}

}

type ConditionStatus string

var ConditionState = struct {
	Succeed string
	Failed  string
}{"Succeed", "Failed"}

type Condition struct {
	Name               string      `json:"name" protobuf:"bytes,1,opt,name=type"`
	Status             string      `json:"status" protobuf:"bytes,2,opt,name=status"`
	ObservedGeneration int64       `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	Reason             string      `json:"reason" protobuf:"bytes,5,opt,name=reason"`
	Message            string      `json:"message" protobuf:"bytes,6,opt,name=message"`
}

func (c *Condition) SetLastTransitionTime(t metav1.Time) *Condition {
	c.LastTransitionTime = t
	return c
}
func (c *Condition) SetSucceed() *Condition {
	c.Status = ConditionState.Succeed
	return c
}
func (c *Condition) SetFailed() *Condition {
	c.Status = ConditionState.Failed
	return c
}
func (c *Condition) SetReason(r string) *Condition {
	c.Reason = r
	return c
}
func (c *Condition) SetMessage(r string) *Condition {
	c.Message = r
	return c
}

func (in *OperatorStatus) setState(state OperatorState) {
	in.State = state
}

func (in *OperatorStatus) SetState(isReady, isHealthy bool) {
	if isHealthy {
		in.setState(OperatorStates.Health)
		return
	}
	if isReady {
		in.setState(OperatorStates.Ready)
		return
	}
	in.setState(OperatorStates.NotReady)
	return
}
