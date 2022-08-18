package cmds

import (
	"flag"
	"fmt"
	"github.com/bleuhold/bh/cmd"
)

var PREMISES *cmd.Command

func init() {
	PREMISES = &cmd.Command{
		Name:        "prem",
		Description: "Properties available.",
		FlagSet:     flag.NewFlagSet("prem", flag.ExitOnError),
		Execute:     premExecute,
	}
	// Define the local usage for a flag.
	PREMISES.FlagSet.BoolVar(&help, "h", false, "")
	PREMISES.FlagSet.BoolVar(&help, "help", false, "Display the prem help.")
	PREMISES.FlagSet.BoolVar(&list, "list", false, "List all the premises available.")
}

func premExecute(cmd *cmd.Command) {
	switch {
	case help:
		cmd.PrintHelp()
	case list:
		fmt.Println("info list")
	default:
		cmd.PrintHelp()
	}
}
