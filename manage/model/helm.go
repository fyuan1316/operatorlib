package model

import "fmt"

type ChartRelease struct {
	ReleaseName string
	Namespace   string
	Values      map[string]interface{}
}

func GetReleaseName(cluster string, chart string) string {
	return fmt.Sprintf("asm-%s-%s", cluster, chart)
}