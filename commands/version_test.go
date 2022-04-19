package commands

import (
	"reflect"
	"testing"

	cli "github.com/tejzpr/zcli/v2"
)

func Test_getVersion(t *testing.T) {
	tests := []struct {
		name string
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
