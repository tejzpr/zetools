package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"zetools/utils"

	"github.com/urfave/cli/v2"
)

const appVersion = "v0.0.1"

func main() {
	app := &cli.App{
		Usage: "Commandline tools for ze common tasks",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print the version of zetools",
				Action: func(c *cli.Context) error {
					fmt.Println(appVersion)
					return nil
				},
			},
			{
				Name:    "base64",
				Aliases: []string{"b64"},
				Usage:   "Base64 Utils",
				Subcommands: []*cli.Command{
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
				},
			},
			{
				Name:    "hmac",
				Aliases: []string{"hm"},
				Usage:   "HMAC Utils",
				Subcommands: []*cli.Command{
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
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
