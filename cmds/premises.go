package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
)

//import (
//	"encoding/json"
//	"flag"
//	"fmt"
//	"github.com/bleuhold/bh/cmd"
//	"github.com/bleuhold/bh/filesys"
//	"github.com/dottics/cli"
//	"github.com/google/uuid"
//	"log"
//	"strings"
//)

var PREMISES *cli.Command
var premisesFilename = "properties.json"

//func init() {
//	PREMISES = &cli.Command{
//		Name:        "prem",
//		Description: "Properties available.",
//		FlagSet:     flag.NewFlagSet("prem", flag.ExitOnError),
//		Execute:     premExecute,
//	}
//	// Define the local usage for a flag.
//	PREMISES.FlagSet.BoolVar(&help, "h", false, "")
//	PREMISES.FlagSet.BoolVar(&help, "help", false, "Display the prem help.")
//	PREMISES.FlagSet.BoolVar(&list, "list", false, "List all the premises available.")
//	PREMISES.FlagSet.BoolVar(&add, "add", false, "To add a new premises.")
//	PREMISES.FlagSet.BoolVar(&update, "update", false, "To update a specific premises' data.")
//	PREMISES.FlagSet.BoolVar(&remove, "remove", false, "To remove/delete a specific premises' data.")
//	PREMISES.FlagSet.StringVar(&name, "name", "", "The name of the premises.")
//	PREMISES.FlagSet.StringVar(&address, "address", "", "The address of the premises.")
//	PREMISES.FlagSet.StringVar(&plotNumber, "plot-number", "", "The plot number of the premises.")
//}

func PremisesExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case B1:
		return ListPremises()
		//case add:
		//	AddPremises(cmd, &name, &address, &plotNumber)
		//default:
		//	cmd.PrintHelp()
	}
	return nil
}

type Premises struct {
	UUID       uuid.UUID `json:"UUID"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	PlotNumber string    `json:"plotNumber"`
}

func (p *Premises) String() string {
	return fmt.Sprintf("%36s %-30s %-8s %s\n", p.UUID, p.Name, p.PlotNumber, p.Address)
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
	fmt.Printf("%-36s %-30s %-8s %s\n", "UUID", "NAME", "PLOTNUM", "ADDRESS")
	for _, pi := range p.Premises {
		fmt.Printf("%s", pi.String())
	}
}

// ListPremises lists all the properties available
func ListPremises() error {
	p := LoadProperties()
	p.Print()
	return nil
}

//// AddPremises adds a new premises.
//func AddPremises(cmd *cmd.Command, name *string, address *string, plotNumber *string) {
//	s := append([]string{*address}, cmd.FlagSet.Args()...)
//	*address = strings.Join(s, " ")
//
//	p := LoadProperties()
//	prem := Premises{
//		UUID:       uuid.New(),
//		Name:       *name,
//		Address:    *address,
//		PlotNumber: *plotNumber,
//	}
//	p.Premises = append(p.Premises, prem)
//	p.Save()
//}
