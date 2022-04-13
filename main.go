package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"

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
							fmt.Println(sEnc)
							os.Exit(0)
							return nil
						},
					},
					{
						Name:  "decode",
						Usage: "Decode a base64 encoded string",
						Action: func(c *cli.Context) error {
							sDec, _ := b64.StdEncoding.DecodeString(c.Args().First())
							fmt.Println(string(sDec))
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
