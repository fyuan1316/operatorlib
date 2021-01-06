package resource

import (
	"fmt"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Order string

var ResourceOrder = struct {
	ASC  Order
	DESC Order
}{"ASC", "DESC"}

type Option func(spec *ResSpec)
type ResSpec struct {
	Order  Order
	Suffix string
}

func Asc() Option {
	return func(spec *ResSpec) {
		spec.Order = ResourceOrder.ASC
	}
}
func Desc() Option {
	return func(spec *ResSpec) {
		spec.Order = ResourceOrder.DESC
	}
}
func Suffix(s string) Option {
	return func(spec *ResSpec) {
		spec.Suffix = s
	}
}

func GetFilesInFolder(folderPath string, opts ...Option) (map[string]string, error) {
	resSpec := &ResSpec{Order: ResourceOrder.ASC}
	for _, opt := range opts {
		opt(resSpec)
	}
	var files []string
	var namedFiles = make(map[string]string)
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if resSpec.Suffix != "" {
			ext := filepath.Ext(path)
			if !strings.EqualFold(ext, resSpec.Suffix) {
				return nil
			}
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.SliceStable(files, func(i, j int) bool {
		var cmp bool
		if files[i] < files[j] {
			cmp = true
		}
		if resSpec.Order == ResourceOrder.ASC {
			return cmp
		} else {
			return !cmp
		}
	})
	for _, path := range files {
		if bytes, err := ioutil.ReadFile(path); err != nil {
			return nil, err
		} else {
			namedFiles[path] = string(bytes)
		}
	}

	return namedFiles, nil
}

func GetChartResources(folderPath string, userValues map[string]interface{}) (map[string]string, error) {
	var err error
	helmChartDirectory := folderPath
	//valuesFilePath := helmChartDirectory + "/values.yaml"

	refChart, err := loader.LoadDir(helmChartDirectory)

	crds := refChart.CRDObjects()
	fmt.Println(crds)
	if err != nil {
		return nil, err
	}
	//var bytes []byte
	var values chartutil.Values
	//if bytes, err = ioutil.ReadFile(valuesFilePath); err != nil {
	//	return nil, err
	//}
	//if values, err = chartutil.ReadValues(bytes); err != nil {
	//	return nil, err
	//}
	if len(userValues) > 0 {
		fmt.Println("override values")
		if values, err = chartutil.CoalesceValues(refChart, userValues); err != nil {
			return nil, err
		}
	}

	isUpgrade := false
	options := chartutil.ReleaseOptions{
		Name:      "asm-operator-test",
		Namespace: "default",
		Revision:  1,
		IsInstall: !isUpgrade,
		IsUpgrade: isUpgrade,
	}

	valuesToRender, err := chartutil.ToRenderValues(refChart, values, options, nil)
	if err != nil {
		return nil, err
	}

	var files map[string]string
	if files, err = engine.Render(refChart, valuesToRender); err != nil {
		return nil, err
	}
	resSep := "---"
	for filePath, content := range files {
		if !hasManifestExtension(filePath) {
			delete(files, filePath)
			continue
		}
		if strings.Contains(content, resSep) {
			resInFile := strings.Split(content, resSep)
			var key string
			for i := range resInFile {
				key = fmt.Sprintf("%s_%d", filePath, i)
				files[key] = resInFile[i]
			}
			delete(files, filePath)
		}
	}
	return files, nil
}
func hasManifestExtension(fname string) bool {
	ext := filepath.Ext(fname)
	return strings.EqualFold(ext, ".yaml") || strings.EqualFold(ext, ".yml") || strings.EqualFold(ext, ".json")
}
