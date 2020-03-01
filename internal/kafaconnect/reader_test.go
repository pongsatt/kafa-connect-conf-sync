package kafaconnect

import (
	"testing"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/config"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
)

func TestClient_Read(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.Config
		wantErr bool
	}{
		{
			name: "read normal",
			want: []*model.Config{
				{
					Name:       "inventory-connector",
					ConfigHash: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kc := NewClient("")
			got, err := config.Read(kc)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("Client.Read() = %v, want %v", len(got), len(tt.want))
			}
		})
	}
}
