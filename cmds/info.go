package cmds

import (
	"flag"
	"fmt"
	"github.com/bleuhold/bh/cmd"
)

var INFO *cmd.Command

func init() {
	INFO = &cmd.Command{
		Name:        "info",
		Description: "To show the current info.",
		FlagSet:     flag.NewFlagSet("info", flag.ExitOnError),
		Execute:     execute,
	}
	// Define the local usage for a flag.
	INFO.FlagSet.BoolVar(&help, "h", false, "")
	INFO.FlagSet.BoolVar(&help, "help", false, "Display the info help.")
	INFO.FlagSet.BoolVar(&list, "list", false, "List all of the current context information.")
}

// executes the info command
func execute(cmd *cmd.Command) {
	fmt.Println("FLAGS", list, help)
	switch {
	case help:
		cmd.PrintHelp()
	case list:
		fmt.Println("info list")
	default:
		cmd.PrintHelp()
	}
}
