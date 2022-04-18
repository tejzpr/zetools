package commands

import "github.com/urfave/cli/v2"

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

// GetCommand returns the command with the given name
func GetCommand(name CommandName, opts *CommandOpts) *cli.Command {
	switch name {
	case Base64CommandName:
		return getCommandInstance(&base64Command{}, opts)
	case HMACCommandName:
		return getCommandInstance(&hmacCommand{}, opts)
	case CheckPortCommandName:
		return getCommandInstance(&checkPortCommand{}, opts)
	case PingCommandName:
		return getCommandInstance(&pingCommand{}, opts)
	case LSCommandName:
		return getCommandInstance(&lsCommand{}, opts)
	}
	return nil
}

func getCommandInstance(cmd Command, opts *CommandOpts) *cli.Command {
	return &cli.Command{
		Name:            string(cmd.Name()),
		Aliases:         cmd.Aliases(),
		Usage:           cmd.Usage(),
		Subcommands:     cmd.Subcommands(),
		Flags:           cmd.Flags(),
		Action:          cmd.Action,
		SkipFlagParsing: opts != nil && opts.SkipFlagParsing,
	}
}
