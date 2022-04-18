package main

import (
	"fmt"
	"os"
	"zetools/commands"

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
			commands.GetCommand(commands.Base64CommandName, nil),
			commands.GetCommand(commands.HMACCommandName, nil),
			commands.GetCommand(commands.CheckPortCommandName, nil),
			commands.GetCommand(commands.PingCommandName, nil),
			commands.GetCommand(commands.LSCommandName, nil),
			commands.GetCommand(commands.TailCommandName, nil),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
