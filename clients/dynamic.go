package clients

/*
import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

func GetDynamicClient(config *rest.Config, gvk schema.GroupVersionKind) (dynamic.NamespaceableResourceInterface, error) {
	dClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	dc, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, err
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}
	return dClient.Resource(mapping.Resource), nil
	//if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
	//
	//} else {
	//	dClient.Resource(mapping.Resource)
	//}

}
*/
