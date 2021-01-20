package equality

import (
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OperatorStatusSemantic = equality.Semantic

func init() {
	_ = OperatorStatusSemantic.AddFunc(func(a, b metav1.Time) bool {
		return true
	})
}
