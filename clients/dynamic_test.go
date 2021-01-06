package clients

/*
import (
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"testing"
)

func TestGetDynamicClient(t *testing.T) {
	type args struct {
		config *rest.Config
		gvk    schema.GroupVersionKind
	}
	tests := []struct {
		name    string
		args    args
		want    dynamic.NamespaceableResourceInterface
		wantErr bool
	}{
		{
			name: "test-post",
			args: args{
				gvk: schema.GroupVersionKind{
					Group:   "apiextensions.k8s.io",
					Version: "v1",
					Kind:    "CustomResourceDefinition",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := GetConfig()
			if err != nil {
				t.Errorf("GetDynamicClient() error = %v", err)
				return
			}
			got, err := GetDynamicClient(config, tt.args.gvk)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDynamicClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			uns, err := got.List(context.Background(), metav1.ListOptions{})
			for _, un := range uns.Items {
				data, err := json.Marshal(un)
				if err != nil {
					panic(err)
				}
				rawObj := json.RawMessage{}
				err = json.Unmarshal(data, &rawObj)
				if err != nil {
					panic(err)
				}

				fmt.Println(rawObj)
			}
			if err != nil {
				t.Errorf("GetDynamicClient() error = %v", err)
				return
			}
			fmt.Println(len(uns.Items))
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetDynamicClient() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
*/
