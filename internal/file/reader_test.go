package file

import (
	"os"
	"testing"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/config"
)

func TestReadFiles(t *testing.T) {
	currentPaht, _ := os.Getwd()

	type args struct {
		rootPath string
	}

	type expectRet struct {
		len int
	}

	tests := []struct {
		name      string
		args      args
		expectRet expectRet
	}{
		{
			name: "list files normally",
			args: args{
				rootPath: currentPaht + "/tests/files",
			},
			expectRet: expectRet{
				len: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &Reader{
				rootPath: tt.args.rootPath,
			}
			configs, err := config.Read(reader)

			if err != nil {
				t.Fatal(err)
			}

			if len(configs) != tt.expectRet.len {
				t.Errorf("expect %d but %d", tt.expectRet.len, len(configs))
			}

			if len(configs) > 0 {
				firstConfig := configs[0]

				if firstConfig.Name == "" {
					t.Error("Name is empty")
				}

				if firstConfig.ConfigHash == "" {
					t.Error("ConfigHash is empty")
				}
			}

		})
	}
}
