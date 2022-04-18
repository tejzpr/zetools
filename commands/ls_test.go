package commands

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func Test_lsCommand_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *lsCommand
		want CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &lsCommand{}
			if got := b.Name(); got != tt.want {
				t.Errorf("lsCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lsCommand_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *lsCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &lsCommand{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lsCommand.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lsCommand_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *lsCommand
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &lsCommand{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("lsCommand.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lsCommand_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *lsCommand
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &lsCommand{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lsCommand.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lsCommand_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *lsCommand
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &lsCommand{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lsCommand.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lsCommand_Action(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		b       *lsCommand
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &lsCommand{}
			if err := b.Action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("lsCommand.Action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
