package model

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type StatusUpdaterFunc func(obj runtime.Object, client client.Client) func(isReady, isHealthy bool) error
type Option func(spec *OperatorOptions)
type OperatorOptions struct {
	Scheme        *runtime.Scheme
	FinalizerID   string
	StatusUpdater StatusUpdaterFunc
	Recorder      record.EventRecorder
}

func SetFinalizer(id string) Option {
	return func(spec *OperatorOptions) {
		spec.FinalizerID = id
	}
}
func SetScheme(scheme *runtime.Scheme) Option {
	return func(spec *OperatorOptions) {
		spec.Scheme = scheme
	}
}
func SetStatusUpdater(updater StatusUpdaterFunc) Option {
	return func(spec *OperatorOptions) {
		spec.StatusUpdater = updater
	}
}
func SetRecorder(recorder record.EventRecorder) Option {
	return func(spec *OperatorOptions) {
		spec.Recorder = recorder
	}
}
