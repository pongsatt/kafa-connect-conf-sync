package kafaconnect

import (
	"testing"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/config"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/file"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
)

func TestClient_Save(t *testing.T) {

	configs, _ := config.Read(file.NewFileReader("/Users/pongsatt/Desktop/Workspaces/labs/kafka-connect-conf-sync/files"))

	type fields struct {
		baseURL string
	}
	type args struct {
		config *model.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Config
		wantErr bool
	}{
		{
			name: "save normal",
			args: args{
				config: configs[0],
			},
			fields: fields{
				baseURL: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kc := NewClient("")
			_, err := kc.Save(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
