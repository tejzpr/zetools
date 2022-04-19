package tail

import (
	"reflect"
	"testing"
	"zetools/commands"

	cli "github.com/tejzpr/zcli/v2"
)

func Test_tailCommand_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *tailCommand
		want commands.CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &tailCommand{}
			if got := b.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tailCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tailCommand_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *tailCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &tailCommand{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tailCommand.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tailCommand_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *tailCommand
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &tailCommand{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("tailCommand.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tailCommand_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *tailCommand
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &tailCommand{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tailCommand.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tailCommand_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *tailCommand
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &tailCommand{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tailCommand.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tailCommand_Action(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		b       *tailCommand
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &tailCommand{}
			if err := b.Action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("tailCommand.Action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
