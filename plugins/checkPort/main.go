package checkport

import (
	"fmt"
	"zetools/commands"

	"github.com/urfave/cli/v2"
)

func init() {
	commands.RegisterCommand(CheckPortCommandName, &checkPortCommand{}, nil)
}

// CheckPortCommandName returns the name of the command
const CheckPortCommandName commands.CommandName = "checkPort"

// checkPortCommand is the check port command
type checkPortCommand struct {
}

// Name returns the name of the command
func (b *checkPortCommand) Name() commands.CommandName {
	return CheckPortCommandName
}

// Aliases returns the aliases of the command
func (b *checkPortCommand) Aliases() []string {
	return []string{"cpt"}
}

// Usage returns the usage of the command
func (b *checkPortCommand) Usage() string {
	return "Check if a port is in use"
}

// Subcommands returns the subcommands of the command
func (b *checkPortCommand) Subcommands() []*cli.Command {
	return nil
}

// Action returns the action of the command
func (b *checkPortCommand) Action(c *cli.Context) error {
	timeout := c.Duration("timeout")
	host := c.String("host")
	port := c.String("port")
	protocol := c.String("protocol")
	if !(protocol == "udp" || protocol == "tcp") {
		return fmt.Errorf("protocol must be either udp or tcp")
	}
	err := CheckPort(protocol, host, port, timeout)
	if err != nil {
		return fmt.Errorf("%s://%s:%s is not accessible", protocol, host, port)
	}
	_, err = fmt.Printf("%s://%s:%s is open", protocol, host, port)
	if err != nil {
		return err
	}
	return nil
}

// Flags returns the flags of the command
func (b *checkPortCommand) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "protocol",
			Aliases:  []string{"pt"},
			Value:    "tcp",
			Usage:    "Protocol to scan for [tcp / udp]",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "timeout",
			Aliases:  []string{"to"},
			Value:    "2s",
			Usage:    "Connection timeout",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "host",
			Value:    "localhost",
			Usage:    "Host to scan for",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "port",
			Value:    "80",
			Usage:    "Port to check",
			Required: false,
		},
	}
}
