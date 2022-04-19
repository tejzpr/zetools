package commands

import (
	"reflect"
	"testing"

	cli "github.com/tejzpr/zcli/v2"
)

func TestRegisterCommand(t *testing.T) {
	type args struct {
		name    CommandName
		command Command
		opts    *CommandOpts
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
			if err := RegisterCommand(tt.args.name, tt.args.command, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("RegisterCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCommands(t *testing.T) {
	tests := []struct {
		name string
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}
