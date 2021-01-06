package chart

/*
import (
	"errors"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/task/chart/sync"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SyncFunction func(client.Client, model.Object) error

type K8sResourceMapping struct {
	ObjectGenerator func() model.Object
	Sync            SyncFunction
	Deletion        SyncFunction
}

var internalMappings = map[metav1.TypeMeta]*K8sResourceMapping{
	metav1.TypeMeta{
		Kind:       "CustomResourceDefinition",
		APIVersion: "apiextensions.k8s.io/v1",
	}: {
		ObjectGenerator: sync.GeneratorCrd,
		Sync:            sync.FnCrd,
	},
	metav1.TypeMeta{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
	}: {
		ObjectGenerator: sync.GeneratorDeployment,
		Sync:            sync.FnDeployment,
	},
	metav1.TypeMeta{
		Kind:       "Service",
		APIVersion: "v1",
	}: {
		ObjectGenerator: sync.GeneratorService,
		Sync:            sync.FnService,
	},
	metav1.TypeMeta{
		Kind:       "ClusterRoleBinding",
		APIVersion: "rbac.authorization.k8s.io/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorClusterRoleBinding,
		Sync:            sync.FnClusterRoleBinding,
	},
	metav1.TypeMeta{
		Kind:       "ClusterRole",
		APIVersion: "rbac.authorization.k8s.io/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorClusterRole,
		Sync:            sync.FnClusterRole,
	},
	metav1.TypeMeta{
		Kind:       "RoleBinding",
		APIVersion: "rbac.authorization.k8s.io/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorRoleBinding,
		Sync:            sync.FnRoleBinding,
	},
	metav1.TypeMeta{
		Kind:       "ClusterConfig",
		APIVersion: "asm.alauda.io/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorClusterConfig,
		Sync:            sync.FnCreateClusterConfig,
	},
	metav1.TypeMeta{
		Kind:       "CaseMonitor",
		APIVersion: "asm.alauda.io/v1beta2",
	}: {
		ObjectGenerator: sync.GeneratorCaseMonitorv1beta2,
		Sync:            sync.FnCreateCaseMonitorv1beta2,
	},
	metav1.TypeMeta{
		Kind:       "CaseMonitor",
		APIVersion: "asm.alauda.io/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorCaseMonitorv1beta1,
		Sync:            sync.FnCreateCaseMonitorv1beta1,
	},
	metav1.TypeMeta{
		Kind:       "CanaryTemplate",
		APIVersion: "asm.alauda.io/v1alpha1",
	}: {
		ObjectGenerator: sync.GeneratorCanaryTemplate,
		Sync:            sync.FnCanaryTemplate,
	},
	metav1.TypeMeta{
		Kind:       "Ingress",
		APIVersion: "extensions/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorIngress,
		Sync:            sync.FnIngress,
	},
	metav1.TypeMeta{
		Kind:       "ServiceMonitor",
		APIVersion: "monitoring.coreos.com/v1",
	}: {
		ObjectGenerator: sync.GeneratorServiceMonitor,
		Sync:            sync.FnServiceMonitor,
	},
	metav1.TypeMeta{
		Kind:       "PodMonitor",
		APIVersion: "monitoring.coreos.com/v1",
	}: {
		ObjectGenerator: sync.GeneratorPodMonitor,
		Sync:            sync.FnPodMonitor,
	},
	metav1.TypeMeta{
		Kind:       "ServiceAccount",
		APIVersion: "v1",
	}: {
		ObjectGenerator: sync.GeneratorServiceAccount,
		Sync:            sync.FnServiceAccount,
	},
	metav1.TypeMeta{
		Kind:       "ValidatingWebhookConfiguration",
		APIVersion: "admissionregistration.k8s.io/v1beta1",
	}: {
		ObjectGenerator: sync.GeneratorValidatingWebhookConfiguration,
		Sync:            sync.FnValidatingWebhookConfiguration,
	},
}

func GetInternalMappings() map[metav1.TypeMeta]*K8sResourceMapping {
	for k := range internalMappings {
		m := internalMappings[k]
		m.Deletion = sync.FnDelete
	}
	return internalMappings
}

func findResource(unStruct unstructured.Unstructured, userMapping map[metav1.TypeMeta]*K8sResourceMapping) (*SyncResource, error) {
	key := metav1.TypeMeta{
		Kind:       unStruct.GetKind(),
		APIVersion: unStruct.GetAPIVersion(),
	}
	var foundMapping *K8sResourceMapping
	if len(userMapping) > 0 {
		if v, ok := userMapping[key]; ok {
			foundMapping = v
		}
	}
	if foundMapping == nil {
		if v, ok := GetInternalMappings()[key]; ok {
			foundMapping = v
		}
	}
	if foundMapping == nil {
		return nil, errors.New(fmt.Sprintf("NotFound type %s in mappings", key))
	}
	return NewSyncResource(foundMapping), nil
}
*/
