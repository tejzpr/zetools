package commands

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func Test_base64Command_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *base64Command
		want CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base64Command{}
			if got := b.Name(); got != tt.want {
				t.Errorf("base64Command.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Command_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *base64Command
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base64Command{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64Command.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Command_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *base64Command
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base64Command{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("base64Command.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Command_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *base64Command
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base64Command{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64Command.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Command_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *base64Command
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base64Command{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64Command.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Command_Action(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		b       *base64Command
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base64Command{}
			if err := b.Action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("base64Command.Action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
