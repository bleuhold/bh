package cmds

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bleuhold/bh/cmd"
	"github.com/bleuhold/bh/filesys"
	"github.com/google/uuid"
	"log"
)

var PREMISES *cmd.Command
var premisesFilename = "properties.json"

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
	PREMISES.FlagSet.BoolVar(&add, "add", false, "To add a new premises.")
	PREMISES.FlagSet.BoolVar(&update, "update", false, "To update a specific premises' data.")
	PREMISES.FlagSet.BoolVar(&remove, "remove", false, "To remove/delete a specific premises' data.")
}

func premExecute(cmd *cmd.Command) {
	switch {
	case help:
		cmd.PrintHelp()
	case list:
		ListProperties()
	default:
		cmd.PrintHelp()
	}
}

type Premises struct {
	UUID       uuid.UUID `json:"UUID"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	PlotNumber string    `json:"plotNumber"`
}

type PremisesData struct {
	Premises []Premises `json:"Premises"`
}

// LoadProperties loads all the property data.
func LoadProperties() *PremisesData {
	p := &PremisesData{
		Premises: []Premises{},
	}
	filesys.LoadInterface(premisesFilename, p)
	return p
}

// Save writes the premises' data to the file system.
func (p *PremisesData) Save() {
	xb, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Unable to marshal premises data: %v", err)
	}
	filesys.WriteFile(premisesFilename, xb)
}

// Print prints the list of all premises
func (p *PremisesData) Print() {
	fmt.Printf("%3s %-13s %-15s %-8s %s\n", "IDX", "UUID", "NAME", "PLOTNUM", "ADDRESS")
	for i, pi := range p.Premises {
		name := pi.Name
		if len(name) > 10 {
			name = name[:12] + "..."
		}
		UUID := pi.UUID.String()
		UUID = UUID[:13]
		fmt.Printf("%-3d %-13s %-15s %-8s %s\n", i, UUID, name, pi.PlotNumber, pi.Address)
	}
}

// ListProperties lists all the properties available
func ListProperties(_ ...string) {
	p := LoadProperties()
	p.Print()
}
