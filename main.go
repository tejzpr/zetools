package main

import (
	"fmt"
	"os"

	"github.com/tejzpr/zetools/commands"

	_ "github.com/tejzpr/zetools-base64"
	_ "github.com/tejzpr/zetools/plugins/checkport"
	_ "github.com/tejzpr/zetools/plugins/hmac"
	_ "github.com/tejzpr/zetools/plugins/ls"
	_ "github.com/tejzpr/zetools/plugins/ping"
	_ "github.com/tejzpr/zetools/plugins/tail"

	cli "github.com/tejzpr/zcli/v2"
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
