package commands

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func Test_hmacCommand_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *hmacCommand
		want CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hmacCommand{}
			if got := b.Name(); got != tt.want {
				t.Errorf("hmacCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hmacCommand_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *hmacCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hmacCommand{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hmacCommand.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hmacCommand_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *hmacCommand
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hmacCommand{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("hmacCommand.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hmacCommand_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *hmacCommand
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hmacCommand{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hmacCommand.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hmacCommand_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *hmacCommand
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hmacCommand{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hmacCommand.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hmacCommand_Action(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		b       *hmacCommand
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &hmacCommand{}
			if err := b.Action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("hmacCommand.Action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
