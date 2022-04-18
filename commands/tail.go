package commands

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/nxadm/tail"
	"github.com/urfave/cli/v2"
)

// TailCommandName returns the name of the command
const TailCommandName CommandName = "tail"

// tailCommand is the base64 command
type tailCommand struct {
}

// Name returns the name of the command
func (b *tailCommand) Name() CommandName {
	return TailCommandName
}

// Aliases returns the aliases of the command
func (b *tailCommand) Aliases() []string {
	return []string{"tl"}
}

// Usage returns the usage of the command
func (b *tailCommand) Usage() string {
	return "Tail files"
}

// Subcommands returns the subcommands of the command
func (b *tailCommand) Subcommands() []*cli.Command {
	return []*cli.Command{}
}

// Flags returns the flags of the command
func (b *tailCommand) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Value:    "",
			Usage:    "The file to tail",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "follow",
			Aliases:  []string{"fw"},
			Value:    "false",
			Usage:    "Follow the file",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "reopen",
			Aliases:  []string{"ro"},
			Value:    "false",
			Usage:    "Reopen the file",
			Required: false,
		},
	}
}

// Action returns the action of the command
func (b *tailCommand) Action(c *cli.Context) error {
	tfile := c.String("file")
	follow := c.Bool("follow")
	reopen := c.Bool("reopen")
	t, err := tail.TailFile(
		tfile, tail.Config{Follow: follow, ReOpen: reopen, MustExist: true})
	if err != nil {
		return err
	}
	cos := make(chan os.Signal, 1)
	signal.Notify(cos, os.Interrupt)
	go func() {
		for _ = range cos {
			t.Stop()
		}
	}()
	// Print the text of each received line
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
	return nil
}
