package commands

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestGetCommand(t *testing.T) {
	type args struct {
		name CommandName
		opts *CommandOpts
	}
	tests := []struct {
		name string
		args args
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommand(tt.args.name, tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCommandInstance(t *testing.T) {
	type args struct {
		cmd  Command
		opts *CommandOpts
	}
	tests := []struct {
		name string
		args args
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommandInstance(tt.args.cmd, tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommandInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
