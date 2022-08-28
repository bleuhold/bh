package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
	"time"
)

var INFO *cli.Command
var infoFilename = "info.json"

//InfoExecute executes the info command.
func InfoExecute(cmd *cli.Command) error {
	switch {
	case B1:
		ListInfo()
	default:
		cmd.PrintHelp()
	}
	return nil
}

type Info struct {
	PropertyUUID uuid.UUID `json:"propertyUUID"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}

// LoadInfo loads the info data.
func LoadInfo() *Info {
	info := &Info{
		PropertyUUID: uuid.UUID{},
		StartDate:    time.Now(),
		EndDate:      time.Now(),
	}
	filesys.LoadInterface(infoFilename, info)
	return info
}

// Print prints the info data to the CLI.
func (i *Info) Print() string {
	s := ""
	s += fmt.Sprintf("%-15s %v\n", "Property UUID:", i.PropertyUUID)
	s += fmt.Sprintf("%-15s %v\n", "Start Date:", i.StartDate.Format("2006-01-02"))
	s += fmt.Sprintf("%-15s %v\n", "End Date:", i.EndDate.Format("2006-01-02"))
	return s
}

// Save writes the data to the file system.
func (i *Info) Save() {
	xb, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Unable to marshal info data: %v", err)
	}
	filesys.WriteFile(infoFilename, xb)
}

// ListInfo lists the info to the console.
func ListInfo() {
	i := LoadInfo()
	fmt.Printf("%s\n", i.Print())
}
