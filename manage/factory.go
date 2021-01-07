package manage

import (
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//var (
//	logger = logging.RegisterScope("controller.oprlib")
//)

type OperatorManage struct {
	K8sClient client.Client
	//CR        Object
	Options  *OperatorOptions
	Recorder record.EventRecorder
}

func NewOperatorManage(client client.Client, opts ...Option) *OperatorManage {
	oprOpts := &OperatorOptions{}
	managerSpec := &OperatorManage{
		K8sClient: client,
		//CR:        cr,
		Options: oprOpts,
	}
	for _, opt := range opts {
		opt(oprOpts)
	}
	return managerSpec
}
