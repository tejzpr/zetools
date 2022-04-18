package commands

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func Test_pingCommand_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *pingCommand
		want CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pingCommand{}
			if got := b.Name(); got != tt.want {
				t.Errorf("pingCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingCommand_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *pingCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pingCommand{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pingCommand.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingCommand_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *pingCommand
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pingCommand{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("pingCommand.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingCommand_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *pingCommand
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pingCommand{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pingCommand.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingCommand_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *pingCommand
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pingCommand{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pingCommand.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingCommand_Action(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		b       *pingCommand
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &pingCommand{}
			if err := b.Action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("pingCommand.Action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
