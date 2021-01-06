package chart

/*
import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"sigs.k8s.io/yaml"
	"text/template"
)

func toYaml(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		fmt.Printf("ToYaml err:%q", err)
		return ""
	}
	return string(data)
}

func Parse(tpl string, values map[string]interface{}) (string, error) {
	fm := sprig.GenericFuncMap()
	fm["toYaml"] = toYaml
	tmpl, err := template.New("key").Funcs(fm).Parse(string(tpl))
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, values)
	return buf.String(), nil
}
*/
