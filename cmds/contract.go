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

var CONTRACT *cli.Command
var contractFilename = "contracts.json"

type Contract struct {
	UUID                 uuid.UUID `json:"uuid"`
	Reference            string    `json:"reference"`
	AdditionalReferences []string  `json:"additionalReferences"`
	Premises             Premises  `json:"premises"`
	Agent                struct {
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		ContactNumber string `json:"contactNumber"`
		Email         string `json:"email"`
	} `json:"agent"`
	Dates struct {
		IncomingInspection time.Time `json:"incomingInspection"`
		OutgoingInspection time.Time `json:"outgoingInspection"`
		Occupation         time.Time `json:"occupation"`
		Evacuation         time.Time `json:"evacuation"`
		Commencement       time.Time `json:"commencement"`
		Termination        time.Time `json:"termination"`
	} `json:"dates"`
	Tenants Tenants `json:"tenants"`
	Fees    struct {
		Agent            float32 `json:"agent"`
		AgentAdditional  float32 `json:"agentAdditional"`
		Deposit          float32 `json:"deposit"`
		DepositIncrease  float32 `json:"depositIncrease"`
		InitialAgreement float32 `json:"initialAgreement"`
		RenewalAgreement float32 `json:"renewalAgreement"`
		Transgression    float32 `json:"transgression"`
	} `json:"fees"`
	Percentages struct {
		Escalation float32 `json:"escalation"`
	} `json:"percentages"`
	Account Account `json:"account"`
}

func (c *Contract) References(tags map[string]bool) map[string]bool {
	tags[c.Reference] = true
	for _, ref := range c.AdditionalReferences {
		tags[ref] = true
	}
	return tags
}

func (c *Contract) String(all bool) string {
	occupation := c.Dates.Occupation.Format("2006-01-02")
	termination := c.Dates.Termination.Format("2006-01-02")
	//unicodeCircle := fmt.Sprintf("\u25EF")
	unicodeCircle := fmt.Sprintf("\u25cf")
	now := time.Now()
	active := ""
	switch {
	case now.Before(c.Dates.Commencement):
		// blue
		active = fmt.Sprintf("\u001b[34m%s\u001B[0m", unicodeCircle)
	case now.After(c.Dates.Termination):
		// red
		active = fmt.Sprintf("\u001b[31m%s\u001B[0m", unicodeCircle)
	case now.Before(c.Dates.Termination) && now.After(c.Dates.Commencement):
		// green
		active = fmt.Sprintf("\u001b[32m%s\u001B[0m", unicodeCircle)
	default:
		//yellow
		active = fmt.Sprintf("\u001b[33m%s\u001B[0m", unicodeCircle)
	}

	s := fmt.Sprintf("%s %36s %-11s %-11s %s\n", active, c.UUID, occupation, termination, c.Premises.Address)
	if all {
		// add tenants to the string
		s += "\n"
		s += c.Tenants.String()
	}
	return s
}

func (c *Contract) Print(all bool) {
	s := fmt.Sprintf("  %-36s %-11s %-11s %s\n", "UUID", "OCCUPATION", "TERMINATION", "ADDRESS")
	s += c.String(all)
	fmt.Println(s)
}

type Contracts map[uuid.UUID]Contract

func LoadContracts() *Contracts {
	xc := &Contracts{}
	filesys.LoadInterface(contractFilename, xc)
	return xc
}

func (xc *Contracts) Save() {
	xb, err := json.Marshal(xc)
	if err != nil {
		log.Fatalf("unable to marshal contracts: %v", err)
	}
	filesys.WriteFile(contractFilename, xb)
}

func (xc *Contracts) String() string {
	s := fmt.Sprintf("  %-36s %-11s %-11s %s\n", "UUID", "OCCUPATION", "TERMINATION", "ADDRESS")
	for _, c := range *xc {
		s += c.String(false)
	}
	return s
}

func (xc *Contracts) Print() {
	fmt.Println(xc.String())
}

/*
	CONTRACT COMMAND
*/

func ContractExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case B1:
		return ListContracts()
	case S1 != "":
		UUID, err := uuid.Parse(S1)
		if err != nil {
			return err
		}
		return ListContract(UUID)
	}
	return nil
}

func ListContracts() error {
	xc := LoadContracts()
	xc.Print()
	return nil
}

func ListContract(UUID uuid.UUID) error {
	xc := LoadContracts()
	c := (*xc)[UUID]
	c.Print(true)
	return nil
}

func GetContract(UUID uuid.UUID) *Contract {
	xc := LoadContracts()
	if c, ok := (*xc)[UUID]; ok {
		return &c
	}
	return nil
}
