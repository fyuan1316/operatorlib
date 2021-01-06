package clients

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func GetConfig() (*rest.Config, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		cfg, err = clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	}
	return cfg, err
}
