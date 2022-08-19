package cmds

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bleuhold/bh/cmd"
	"github.com/bleuhold/bhman/fs"
	"github.com/google/uuid"
	"log"
	"time"
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
	INFO.FlagSet.BoolVar(&set, "set", false, "Set some info parameter used within the global context of this application.")
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

type Info struct {
	PropertyUUID uuid.UUID `json:"propertyUUID"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}

// Print prints the info data to the CLI.
func (i *Info) Print() {
	fmt.Printf("%-15s %v\n", "Property UUID:", i.PropertyUUID)
	fmt.Printf("%-15s %v\n", "Start Date:", i.StartDate)
	fmt.Printf("%-15s %v\n", "End Date:", i.EndDate)
}

// Save writes the data to the file system.
func (i *Info) Save() {
	xb, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Unable to marshal info data: %v", err)
	}
	fs.WriteFile("info.json", xb)
}
