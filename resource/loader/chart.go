package loader

/*
import (
	"helm.sh/helm/v3/pkg/task/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"io/ioutil"
)

func LoadChart(dir string) (map[string]string, error) {
	//TODO fy
	dir = "/Users/yuan/Dev/work/GolangProjects/charts/cluster-asm-cluster-asm-copy/cluster-asm"
	valuesFilePath := dir + "/values.yaml"
	var err error
	refChart, err := loader.LoadDir(dir)
	if err != nil {
		return nil, err
	}
	var bytes []byte
	var values chartutil.Values
	if bytes, err = ioutil.ReadFile(valuesFilePath); err != nil {
		return nil, err
	}
	isUpgrade := false
	options := chartutil.ReleaseOptions{
		Name:      "asm-operator-test",
		Namespace: "default",
		Revision:  1,
		IsInstall: !isUpgrade,
		IsUpgrade: isUpgrade,
	}

	if values, err = chartutil.ReadValues(bytes); err != nil {
		return nil, err
	}
	valuesToRender, err := chartutil.ToRenderValues(refChart, values, options, nil)
	if err != nil {
		return nil, err
	}
	var files map[string]string
	if files, err = engine.Render(refChart, valuesToRender); err != nil {
		return nil, err
	}
	return files, nil

}
*/
