package chart

import (
	"fmt"
	"github.com/fyuan1316/klient"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"path/filepath"
	"strings"
)

var _ model.ExecuteItem = ChartTask2{}

//var _ model.Chart = &ChartTask2{}

type ChartTask2 struct {
	//子类override 接口
	implementor model.OverrideOperation
	// task path
	Dir    string
	Chart  *chart.Chart
	values chartutil.Values
	klient *klient.Client
}

func (c *ChartTask2) Override(operation model.OverrideOperation) {
	c.implementor = operation
}
func (c *ChartTask2) Init() {
	var (
		chart  *chart.Chart
		client *klient.Client
		err    error
	)
	if chart, err = loader.LoadDir(c.Dir); err != nil {
		panic(err)
	}
	if client, err = klient.NewE("", ""); err != nil {
		panic(err)
	}
	c.klient = client
	c.Chart = chart

}
func (c *ChartTask2) apply(userValues map[string]interface{}) error {
	var err error
	if err = c.applyCrds(); err != nil {
		return err
	}
	//set values
	resFiles, err := c.generateFiles(userValues)
	if err != nil {
		return err
	}
	for _, crd := range resFiles {
		if err = c.klient.Apply([]byte(crd)); err != nil {
			return err
		}
	}
	return nil
}
func (c *ChartTask2) delete(userValues map[string]interface{}) error {
	var err error

	//set values
	resFiles, err := c.generateFiles(userValues)
	if err != nil {
		return err
	}
	for _, crd := range resFiles {
		if err = c.klient.Delete([]byte(crd)); err != nil {
			return err
		}
	}
	return nil
}

func (c *ChartTask2) generateFiles(userValues map[string]interface{}) (Files, error) {
	var err error
	c.values = c.Chart.Values
	// merge values if any
	if len(userValues) > 0 {
		fmt.Println("override values")
		if c.values, err = chartutil.CoalesceValues(c.Chart, userValues); err != nil {
			return nil, err
		}
	}
	//render files
	isUpgrade := false
	options := chartutil.ReleaseOptions{
		Name:      "asm-operator-test",
		Namespace: "default",
		Revision:  1,
		IsInstall: !isUpgrade,
		IsUpgrade: isUpgrade,
	}
	valuesToRender, err := chartutil.ToRenderValues(c.Chart, c.values, options, nil)
	if err != nil {
		return nil, err
	}
	var files Files
	if files, err = engine.Render(c.Chart, valuesToRender); err != nil {
		return nil, err
	}
	// filter files
	resFiles := files.filterManifestExtension(func(filename string) bool {
		ext := filepath.Ext(filename)
		return strings.EqualFold(ext, ".yaml") || strings.EqualFold(ext, ".yml") || strings.EqualFold(ext, ".json")
	}).splitOneResourcePerFile()
	return resFiles, nil
}

type Files map[string]string

func (files Files) filterManifestExtension(fn func(filename string) bool) Files {
	m := make(map[string]string, 0)
	for k, v := range files {
		if fn(k) {
			m[k] = v
		}
	}
	return m
}
func (files Files) splitOneResourcePerFile() Files {
	resSep := "---"
	m := make(map[string]string, 0)
	for filePath, content := range files {
		if strings.Contains(content, resSep) {
			resInFile := strings.Split(content, resSep)
			var key string
			for i := range resInFile {
				key = fmt.Sprintf("%s_%d", filePath, i)
				m[key] = resInFile[i]
			}
		} else {
			m[filePath] = content
		}
	}
	return m
}

func (c *ChartTask2) applyCrds() error {
	if len(c.Chart.CRDObjects()) > 0 {
		for _, crd := range c.Chart.CRDObjects() {
			if err := c.klient.Apply(crd.File.Data); err != nil {
				return err
			}
		}
	}
	return nil
}

//func (c *ChartTask2) Reload(userValues map[string]interface{}) (err error) {
//	if len(userValues) > 0 {
//		fmt.Println("override values") //TODO fy
//		var values chartutil.Values
//		if values, err = chartutil.CoalesceValues(c.Chart, userValues); err != nil {
//			return err
//		}
//		c.values = values
//	}
//	return nil
//}

func (c ChartTask2) Name() string {
	panic("implement me")
}
func (c ChartTask2) GetImplementor() model.OverrideOperation {
	return c.implementor
}
func (c ChartTask2) Run(ctx *model.OperatorContext) error {
	if c.GetImplementor().GetOperation() == model.Operations.Provision {
		return c.Apply(ctx)
	} else if c.GetImplementor().GetOperation() == model.Operations.Deletion {
		return c.Delete(ctx)
	} else {
		return errors.New("UnSupport type of ResourceTask")
	}
}
func (c ChartTask2) Apply(ctx *model.OperatorContext) error {
	return c.apply(ctx.OperatorParams)
}
func (c ChartTask2) Delete(ctx *model.OperatorContext) error {
	return c.delete(ctx.OperatorParams)
}
