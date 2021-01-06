package chart

/*
import (
	"fmt"
	"testing"
)

func TestSyncManager_LoadFile(t *testing.T) {
	type fields struct {
		K8sResource SyncResource
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test-load-file",
			fields:  fields{},
			args:    args{filePath: "./pod.yaml"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ChartTask{
				FileTask: FileTask{
					K8sResource: map[string]SyncResource{
						"test1": tt.fields.K8sResource,
					},
				},
			}

			if err := m.LoadFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println()
		})
	}
}
*/
