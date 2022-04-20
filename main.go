package main

import (
	"fmt"
	"os"

	"github.com/tejzpr/zetools/commands"

	// Available command plugins
	_ "github.com/tejzpr/zetools-base64"
	_ "github.com/tejzpr/zetools-cat"
	_ "github.com/tejzpr/zetools-checkport"
	_ "github.com/tejzpr/zetools-hmac"
	_ "github.com/tejzpr/zetools-ls"
	_ "github.com/tejzpr/zetools-ping"
	_ "github.com/tejzpr/zetools-tail"

	// End of available command plugins

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
