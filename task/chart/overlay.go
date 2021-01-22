package chart

import (
	"fmt"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart/object"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"strings"
)

func (c *ChartTask) generateOverlayFiles(ctx *model.OperatorContext) (Files, error) {
	var err error
	labels, err := getCoreOwnerLabels(ctx)
	if err != nil {
		return nil, err
	}
	namedFiles, err := c.generateFiles(ctx.InstalledNamespace, ctx.OperatorParams)
	if err == nil && len(labels) > 0 {
		var ymlSlice []string
		for _, yml := range namedFiles {
			ymlSlice = append(ymlSlice, yml)
		}
		k8sObjects, err := object.ParseK8sObjectsFromYAMLManifest(strings.Join(ymlSlice, object.YAMLSeparator))
		if err != nil {
			return nil, err
		}
		labeledFiles := make(map[string]string)
		for i, obj := range k8sObjects {
			for k, v := range labels {
				err := SetLabel(obj.UnstructuredObject(), k, v)
				if err != nil {
					return nil, err
				}
			}
			bytes, err := obj.ParsedYAML()
			if err != nil {
				return nil, err
			}
			labeledFiles[string(i)] = string(bytes)
		}
		return labeledFiles, nil
	}
	return namedFiles, err
}

func getCRName(ctx *model.OperatorContext) (string, error) {
	objAccessor, err := meta.Accessor(ctx.Instance)
	if err != nil {
		return "", err
	}
	return objAccessor.GetName(), nil
}
func getCRNamespace(ctx *model.OperatorContext) (string, error) {
	objAccessor, err := meta.Accessor(ctx.Instance)
	if err != nil {
		return "", err
	}
	return objAccessor.GetNamespace(), nil
}

func getOwningResourcePrefix(ctx *model.OperatorContext) string {
	k := ctx.Instance.GetObjectKind().GroupVersionKind().Kind
	g := ctx.Instance.GetObjectKind().GroupVersionKind().Group
	return strings.ToLower(fmt.Sprintf("%s.%s", k, g))
}
func getOwningResourceName(prefix string) string {
	return fmt.Sprintf("%s/owning-resource", prefix)
}
func getOwningResourceNamespaceKey(prefix string) string {
	return fmt.Sprintf("%s/owning-resource-namespace", prefix)
}
func getCoreOwnerLabels(ctx *model.OperatorContext) (map[string]string, error) {
	labelKeyPrefix := getOwningResourcePrefix(ctx)
	crName, err := getCRName(ctx)
	if err != nil {
		return nil, err
	}
	crNamespace, err := getCRNamespace(ctx)
	if err != nil {
		return nil, err
	}
	labels := make(map[string]string)

	labels[getOwningResourceName(labelKeyPrefix)] = crName
	labels[getOwningResourceNamespaceKey(labelKeyPrefix)] = crNamespace

	return labels, nil
}

func SetLabel(resource runtime.Object, label, value string) error {
	resourceAccessor, err := meta.Accessor(resource)
	if err != nil {
		return err
	}
	labels := resourceAccessor.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	labels[label] = value
	resourceAccessor.SetLabels(labels)
	return nil
}
