package main

import (
	"fmt"
	"os"
	"zetools/commands"

	_ "zetools/plugins/checkport"
	_ "zetools/plugins/hmac"
	_ "zetools/plugins/ls"
	_ "zetools/plugins/ping"
	_ "zetools/plugins/tail"

	_ "github.com/tejzpr/zetools-base64"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage:    "Unified commandline tools for common tasks",
		Commands: commands.GetCommands(),
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
