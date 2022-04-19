package checkport

import (
	"testing"
	"time"
)

func TestCheckPort(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     string
		timeout  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPort(tt.args.protocol, tt.args.host, tt.args.port, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("CheckPort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
