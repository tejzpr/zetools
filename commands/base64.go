package commands

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/urfave/cli/v2"
)

// Base64CommandName returns the name of the command
const Base64CommandName CommandName = "base64"

// base64Command is the base64 command
type base64Command struct {
}

// Name returns the name of the command
func (b *base64Command) Name() CommandName {
	return Base64CommandName
}

// Aliases returns the aliases of the command
func (b *base64Command) Aliases() []string {
	return []string{"b64"}
}

// Usage returns the usage of the command
func (b *base64Command) Usage() string {
	return "Base64 Utils"
}

// Subcommands returns the subcommands of the command
func (b *base64Command) Subcommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "encode",
			Usage: "Encode a string to Base64",
			Action: func(c *cli.Context) error {
				sEnc := b64.StdEncoding.EncodeToString([]byte(c.Args().First()))
				fmt.Print(sEnc)
				return nil
			},
		},
		{
			Name:  "decode",
			Usage: "Decode a base64 encoded string",
			Action: func(c *cli.Context) error {
				sDec, _ := b64.StdEncoding.DecodeString(c.Args().First())
				fmt.Print(string(sDec))
				return nil
			},
		},
	}
}

// Flags returns the flags of the command
func (b *base64Command) Flags() []cli.Flag {
	return []cli.Flag{}
}

// Action returns the action of the command
func (b *base64Command) Action(c *cli.Context) error {
	return nil
}
