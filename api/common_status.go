package api

type OperationStage string

var OperationStages = struct {
	PreCheck  OperationStage
	PreRun    OperationStage
	Run       OperationStage
	PostRun   OperationStage
	PostCheck OperationStage
}{"PreCheck", "PreRun", "Run", "PostRun", "PostCheck"}

func (in OperatorStatus) GetOperatorStatus() OperatorStatus {
	return in
}
