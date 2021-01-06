package manage

import (
	//"github.com/fyuan1316/asm-operator/pkg/logging"
	"github.com/fyuan1316/operatorlib/manage/model"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//var (
//	logger = logging.RegisterScope("controller.oprlib")
//)

func NewOperatorManage(client client.Client, opts ...model.Option) *model.OperatorManage {
	oprOpts := &model.OperatorOptions{}
	managerSpec := &model.OperatorManage{
		K8sClient: client,
		//CR:        cr,
		Options: oprOpts,
	}
	for _, opt := range opts {
		opt(oprOpts)
	}
	return managerSpec
}
