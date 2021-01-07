package api

// +k8s:deepcopy-gen=false

type OperatorSpec struct {
	Parameters string `json:"parameters,omitempty"`
}

type OperatorState string

var OperatorStates = struct {
	NotReady OperatorState
	Ready    OperatorState
	Health   OperatorState
}{"NotReady", "Ready", "Health"}

type OperatorStatus struct {
	State OperatorState `json:"state,omitempty"`
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
