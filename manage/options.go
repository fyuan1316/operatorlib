package manage

import (
	"github.com/fyuan1316/operatorlib/manage/model"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

type StatusUpdaterFunc func(*model.OperatorContext, *model.StatusContext) error
type Option func(spec *OperatorOptions)
type OperatorOptions struct {
	Scheme        *runtime.Scheme
	FinalizerID   string
	StatusUpdater StatusUpdaterFunc
	Recorder      record.EventRecorder
	//兼容captain安装信息
	DefaultInstallNamespace string
	ChartName               string
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

func SetDefaultInstallNamespace(ns string) Option {
	return func(spec *OperatorOptions) {
		spec.DefaultInstallNamespace = ns
	}
}
func SetChartName(chart string) Option {
	return func(spec *OperatorOptions) {
		spec.ChartName = chart
	}
}
