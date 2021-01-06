package chart

/*
import (
	"fmt"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"io/ioutil"
	"testing"
)

func GetDefaults() map[string]interface{} {
	values := TreeValue(map[string]interface{}{
		"Release.Namespace":                                 "default",
		"Release.Name":                                      "asm-operator",
		"Release.Service":                                   "asm-operator",
		"Values.global.scheme":                              "http",
		"Values.global.host":                                "k8s.alauda.io",
		"Values.global.useNodePort":                         false,
		"Values.global.labelBaseDomain":                     "alauda.io",
		"Values.global.alb2Name":                            "",
		"Values.global.registry.address":                    "harbor.alauda.cn",
		"Values.asm_controller.resources.limits.cpu":        "300m",
		"Values.asm_controller.resources.limits.memory":     "300Mi",
		"Values.asm_controller.resources.requests.cpu":      "300m",
		"Values.asm_controller.resources.requests.memory":   "300Mi",
		"Values.asm_controller.imagepullPolicy":             "IfNotPresent",
		"Values.asm_controller.replicaCount":                1,
		"Values.grafana.url":                                "https://asm.dev.com/grafana-asm-cluster1",
		"Values.grafana.basepath":                           "/grafana-asm",
		"Values.grafana.tls":                                "null",
		"Values.grafana.servicePort":                        "3000",
		"Values.prometheus.url":                             "",
		"Values.prometheus.serviceMonitorLabels.prometheus": "kube-prometheus",
		"Values.istio.namespace":                            "istio-system",
		"Values.istio.traceSampling":                        100,
		"Values.istio.includeIPRanges":                      "*",
		"Values.istio.gatewayNodeportValue":                 30666,
		"Values.istio.gatewayHttpsNodeportValue":            30665,
		"Values.jaeger.url":                                 "https://asm.dev.com/jaeger-asm-cluster1",
		"Values.jaeger.query.basepath":                      "/jaeger-asm",
		"Values.jaeger.query.replicas":                      1,
		"Values.jaeger.collector.replicas":                  1,
		"Values.jaeger.elasticsearch.indexprefix":           "asm",
		"Values.jaeger.elasticsearch.serverurl":             "http://es-elasticsearch-client.istio-system:9200",
		"Values.jaeger.elasticsearch.username":              "elastic",
		"Values.jaeger.elasticsearch.password":              "changeme",
		"Values.jaeger.elasticsearch.indexRetainsDays":      "3",
		"Values.jaeger.ui.port":                             16686,
		"Values.jaeger.nsZipkinCollector.host":              "ns-zipkin",
		"Values.jaeger.nsZipkinCollector.port":              9411,
		"Values.jaeger.ingress.enabled":                     false,
		//"Values.jaeger.ingress.hosts":              "null",
		//"Values.jaeger.ingress.annotations":              null,
		//"Values.jaeger.ingress.tls":              null,
		"Values.jaeger.replicaCount": 1,
		//"Values.jaeger.zipkinService.annotations":              ,
		"Values.jaeger.zipkinService.portName":     "http",
		"Values.jaeger.zipkinService.type":         "ClusterIP",
		"Values.jaeger.zipkinService.externalPort": 9411,
		"Values.jaeger.zipkinService.internalPort": 9411,
		"Values.jaeger.resources.limits.cpu":       "100m",
		"Values.jaeger.resources.limits.memory":    "300Mi",
		"Values.jaeger.resources.requests.cpu":     "100m",
		"Values.jaeger.resources.requests.memory":  "300Mi",
	})

	return values
}

func Test_loadChart(t *testing.T) {
	helmChartDirectory := "/Users/yuan/Dev/GolangProjects/asm-operator/files/provision/cluster-asm"
	valuesFilePath := helmChartDirectory + "/values.yaml"
	var err error
	refChart, err := loader.LoadDir(helmChartDirectory)
	if err != nil {
		panic(err)
	}
	var bytes []byte
	var values chartutil.Values
	if bytes, err = ioutil.ReadFile(valuesFilePath); err != nil {
		panic(err)
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
		panic(err)
	}
	valuesToRender, err := chartutil.ToRenderValues(refChart, values, options, nil)
	if err != nil {
		panic(err)
	}

	if files, err := engine.Render(refChart, valuesToRender); err != nil {
		panic(err)
	} else {
		fmt.Println(len(files))
	}

	fmt.Println()

}
*/
