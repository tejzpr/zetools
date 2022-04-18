package commands

import (
	"fmt"
	"zetools/utils"

	"github.com/urfave/cli/v2"
)

// HMACCommandName returns the name of the command
const HMACCommandName CommandName = "hmac"

// hmacCommand is the base64 command
type hmacCommand struct {
}

// Name returns the name of the command
func (b *hmacCommand) Name() CommandName {
	return HMACCommandName
}

// Aliases returns the aliases of the command
func (b *hmacCommand) Aliases() []string {
	return []string{"hm"}
}

// Usage returns the usage of the command
func (b *hmacCommand) Usage() string {
	return "HMAC Utils"
}

// Subcommands returns the subcommands of the command
func (b *hmacCommand) Subcommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "sha256",
			Aliases: []string{"256"},
			Usage:   "Hash a string with SHA256",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "text",
					Aliases:  []string{"t"},
					Value:    "",
					Usage:    "Text to be hashed",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "filename",
					Aliases:  []string{"f"},
					Value:    "",
					Usage:    "File to be hashed",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "key",
					Aliases:  []string{"k"},
					Value:    "",
					Usage:    "Secret key for hashing",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				text := c.String("text")
				filename := c.String("filename")
				key := c.String("key")
				out, err := utils.HMAC(text, filename, key, "sha256")
				if err != nil {
					return err
				}
				_, err = fmt.Print(string(out))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "sha512",
			Aliases: []string{"512"},
			Usage:   "Hash a string with SHA512",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "text",
					Aliases:  []string{"t"},
					Value:    "",
					Usage:    "Text to be hashed",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "filename",
					Aliases:  []string{"f"},
					Value:    "",
					Usage:    "File to be hashed",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "key",
					Aliases:  []string{"k"},
					Value:    "",
					Usage:    "Secret key for hashing",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				text := c.String("text")
				filename := c.String("filename")
				key := c.String("key")
				out, err := utils.HMAC(text, filename, key, "sha512")
				if err != nil {
					return err
				}
				_, err = fmt.Print(string(out))
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}

// Flags returns the flags of the command
func (b *hmacCommand) Flags() []cli.Flag {
	return []cli.Flag{}
}

// Action returns the action of the command
func (b *hmacCommand) Action(c *cli.Context) error {
	return nil
}
