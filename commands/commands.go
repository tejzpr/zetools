package commands

import (
	"fmt"
	"sync"

	cli "github.com/tejzpr/zcli/v2"
)

// CommandName is the name of the command
type CommandName string

// Command is the interface that all commands must implement
type Command interface {
	Name() CommandName
	Aliases() []string
	Usage() string
	Subcommands() []*cli.Command
	Flags() []cli.Flag
	Action(c *cli.Context) error
}

// CommandOpts is the options for a command
type CommandOpts struct {
	// Treat all flags as normal arguments if true
	SkipFlagParsing bool
}

var (
	commandMu sync.RWMutex
	commands  = make(map[CommandName]*cli.Command)
)

// RegisterCommand registers a command
func RegisterCommand(name CommandName, command Command, opts *CommandOpts) error {
	commandMu.Lock()
	defer commandMu.Unlock()
	if command == nil {
		return fmt.Errorf("command is nil")
	}
	if _, dup := commands[name]; dup {
		return fmt.Errorf("command already registered: %s", name)
	}
	commands[name] = &cli.Command{
		Name:            string(command.Name()),
		Aliases:         command.Aliases(),
		Usage:           command.Usage(),
		Subcommands:     command.Subcommands(),
		Flags:           command.Flags(),
		Action:          command.Action,
		SkipFlagParsing: opts != nil && opts.SkipFlagParsing,
	}
	return nil
}

// GetCommands returns the registered commands
func GetCommands() []*cli.Command {
	commands["version"] = getVersion()
	cmds := make([]*cli.Command, 0)
	for _, command := range commands {
		cmds = append(cmds, command)
	}
	return cmds
}
