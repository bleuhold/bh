package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
)

var ACCOUNT *cli.Command
var accountsFilename = "accounts.json"

type Account struct {
	UUID         uuid.UUID `json:"uuid"`
	Number       string    `json:"number"`
	AccountType  string    `json:"accountType"`
	ProviderName string    `json:"providerName"`
	HolderName   string    `json:"holderName"`
}

// NewAccount creates a new account, with a default UUID.
func NewAccount(number, accountType, providerName, holderName string) *Account {
	a := &Account{
		UUID:         uuid.New(),
		Number:       number,
		AccountType:  accountType,
		ProviderName: providerName,
		HolderName:   holderName,
	}
	return a
}

type Accounts []Account

// LoadAccounts loads the accounts' information from the file system.
func LoadAccounts() *Accounts {
	xa := &Accounts{}
	filesys.LoadInterface(accountsFilename, xa)
	return xa
}

// Save writes the accounts' data to the file system.
func (xa *Accounts) Save() {
	xb, err := json.Marshal(xa)
	if err != nil {
		log.Fatalf("Unable to marshal accounts data: %v", err)
	}
	filesys.WriteFile(accountsFilename, xb)
}

func (xa *Accounts) String() string {
	s := fmt.Sprintf("%-36s %-16s %-16s %-16s\n", "UUID", "NUMBER", "PROVIDER", "TYPE")
	for _, a := range *xa {
		s += fmt.Sprintf("%36s %-16s %-16s %-16s\n", a.UUID, a.Number, a.ProviderName, a.AccountType)
	}
	return s
}

func (xa *Accounts) Add(acc *Account) *Accounts {
	x := *xa
	x = append(x, *acc)
	return &x
}

func GetAccount(UUID uuid.UUID) (*Account, error) {
	xa := LoadAccounts()
	for _, a := range *xa {
		if a.UUID == UUID {
			return &a, nil
		}
	}
	return &Account{}, fmt.Errorf("invalid UUID: account not found: %v", UUID)
}

/*
	ACCOUNT
*/

func AccountExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case B1:
		ListAccounts()
		return nil
	}
	return nil
}

func ListAccounts() {
	xa := LoadAccounts()
	fmt.Print(xa.String())
}
