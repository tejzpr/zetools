package ls

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"zetools/commands"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func init() {
	commands.RegisterCommand(LSCommandName, &lsCommand{}, nil)
}

// LSCommandName returns the name of the command
const LSCommandName commands.CommandName = "ls"

// lsCommand is the base64 command
type lsCommand struct {
}

// Name returns the name of the command
func (b *lsCommand) Name() commands.CommandName {
	return LSCommandName
}

// Aliases returns the aliases of the command
func (b *lsCommand) Aliases() []string {
	return []string{"ls"}
}

// Usage returns the usage of the command
func (b *lsCommand) Usage() string {
	return "List Directory"
}

// Subcommands returns the subcommands of the command
func (b *lsCommand) Subcommands() []*cli.Command {
	return []*cli.Command{}
}

// Flags returns the flags of the command
func (b *lsCommand) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Aliases:  []string{"p"},
			Value:    ".",
			Usage:    "Path to list",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "block-size",
			Aliases:  []string{"bs"},
			Value:    "B",
			Usage:    "Block Size, possible vales are [B (bytes), K (kilo bytes), M (mega bytes), G (giga bytes)",
			Required: false,
		},
	}
}

// Action returns the action of the command
func (b *lsCommand) Action(c *cli.Context) error {
	pathList := c.String("path")
	blockSize := c.String("bs")
	blockType := "B"
	switch blockSize {
	case "B":
		blockType = "Bytes"
	case "K":
		blockType = "KiloBytes"
	case "M":
		blockType = "MegaBytes"
	case "G":
		blockType = "GigaBytes"
	default:
		return fmt.Errorf("Invalid block size %s", blockSize)
	}
	files, err := ioutil.ReadDir(pathList)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return fmt.Errorf("Operation not permitted")
		}
		return fmt.Errorf("Error listing directory")
	}
	data := [][]string{}
	for _, file := range files {
		size := fmt.Sprintf("%d", file.Size())
		if blockSize == "K" {
			size = fmt.Sprintf("%f", float64(file.Size())/(1<<10))
		} else if blockSize == "M" {
			size = fmt.Sprintf("%f", float64(file.Size())/(1<<20))
		} else if blockSize == "G" {
			size = fmt.Sprintf("%f", float64(file.Size())/(1<<30))
		}

		data = append(data, []string{file.Name(), file.Mode().String(), fmt.Sprintf("%t", file.IsDir()), size, file.ModTime().Format(time.ANSIC)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Mode", "Is Dir?", "Size (" + blockType + ")", "Modified Time"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()

	return nil
}
