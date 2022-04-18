package commands

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func Test_checkPortCommand_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *checkPortCommand
		want CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &checkPortCommand{}
			if got := b.Name(); got != tt.want {
				t.Errorf("checkPortCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPortCommand_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *checkPortCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &checkPortCommand{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkPortCommand.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPortCommand_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *checkPortCommand
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &checkPortCommand{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("checkPortCommand.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPortCommand_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *checkPortCommand
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &checkPortCommand{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkPortCommand.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPortCommand_Action(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		b       *checkPortCommand
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &checkPortCommand{}
			if err := b.Action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("checkPortCommand.Action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkPortCommand_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *checkPortCommand
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &checkPortCommand{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkPortCommand.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}
