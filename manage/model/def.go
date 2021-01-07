package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//
//type OperatorManage struct {
//	K8sClient client.Client
//	Options  *manage.OperatorOptions
//	Recorder record.EventRecorder
//}

//func (m *OperatorManage) GetEditableCR() runtime.Object {
//	return m.CR.DeepCopyObject()
//}

type Object interface {
	runtime.Object
	metav1.Object
	//CommonOperator
}

type CommonOperator interface {
	runtime.Object
	metav1.Object
	GetOperatorParams() (map[string]interface{}, error)
	//SetOperatorStatus(api.OperatorStatus)
	//GetOperatorStatus() api.OperatorStatus
}
type OperatorContext struct {
	K8sClient client.Client
	//Options   *OperatorOptions
	Recorder       record.EventRecorder
	Instance       CommonOperator
	OperatorParams map[string]interface{}
}
