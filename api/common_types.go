package api

// +k8s:deepcopy-gen=false

type OperatorSpec struct {
	Parameters map[string]string `json:"parameters,omitempty"`
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
