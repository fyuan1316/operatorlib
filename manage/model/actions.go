package model

type PreRun interface {
	PreRun(*OperatorContext) error
}
type PostRun interface {
	PostRun(*OperatorContext) error
}
type PreCheck interface {
	PreCheck(*OperatorContext) (bool, error)
}
type PostCheck interface {
	PostCheck(*OperatorContext) (bool, error)
}
type Runnable interface {
	Run(*OperatorContext) error
}

type HealthCheck interface {
	IsReady(*OperatorContext) bool
	IsHealthy(*OperatorContext) bool
}

//
func CanDoPreCheck(inf interface{}) (PreCheck, bool) {
	if b, ok := inf.(PreCheck); ok {
		return b, true
	}
	return nil, false
}
func CanDoPostCheck(inf interface{}) (PostCheck, bool) {
	if b, ok := inf.(PostCheck); ok {
		return b, true
	}
	return nil, false
}
func CanDoPreRun(inf interface{}) (PreRun, bool) {
	if b, ok := inf.(PreRun); ok {
		return b, true
	}
	return nil, false
}
func CanDoPostRun(inf interface{}) (PostRun, bool) {
	if b, ok := inf.(PostRun); ok {
		return b, true
	}
	return nil, false
}
func CanDoHealthCheck(inf interface{}) (HealthCheck, bool) {
	if b, ok := inf.(HealthCheck); ok {
		return b, true
	}
	return nil, false
}

//
//func CanDoRun(inf interface{}) (Runnable, bool) {
//	if b, ok := inf.(Runnable); ok {
//		return b, true
//	}
//	return nil, false
//}

//func IsChartTask(inf interface{}) (Chart, bool) {
//	if b, ok := inf.(Chart); ok {
//		return b, true
//	}
//	return nil, false
//}

//type Chart interface {
//	Reload(map[string]interface{}) error
//}
type Item interface {
	Name() string
}
type ExecuteItem interface {
	Item
	Runnable
}
type OperationType string

var Operations = struct {
	Provision OperationType
	Deletion  OperationType
}{"Provision", "Deletion"}

// 子类需要实现的接口集
type OverrideOperation interface {
	GetOperation() OperationType
	Override(OverrideOperation)
}
