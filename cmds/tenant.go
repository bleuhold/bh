package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
)

var TENANT *cli.Command
var tenantFilename = "tenants.json"

type Tenant struct {
	UUID          uuid.UUID `json:"uuid"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	ID            string    `json:"id"`
	Passport      string    `json:"passport"`
	ContactNumber string    `json:"contactNumber"`
	Email         string    `json:"email"`
}

func (t *Tenant) String() string {
	return fmt.Sprintf("%36s %13s %-15s %-15s %-13s %-10s %s\n", t.UUID, t.ID, t.FirstName, t.LastName, t.ContactNumber, t.Passport, t.Email)
}

type Tenants []Tenant

func LoadTenants() *Tenants {
	t := &Tenants{}
	filesys.LoadInterface(tenantFilename, t)
	return t
}

func (t *Tenants) Save() {
	xb, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("unable to marshal tenants data: %v", err)
	}
	filesys.WriteFile(tenantFilename, xb)
}

func (t *Tenants) String() string {
	s := fmt.Sprintf("%-36s %-13s %-15s %-15s %-13s %-10s %s\n", "UUID", "ID", "FIRSTNAME", "LASTNAME", "CONTACTNUMBER", "PASSPORT", "EMAIL")
	for _, ti := range *t {
		s += ti.String()
	}
	return s
}

/*
	COMMAND
*/

func TenantExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case B1:
		ListTenants()
		return nil
	}
	return nil
}

func ListTenants() {
	xt := LoadTenants()
	fmt.Printf("%v", xt)
}
