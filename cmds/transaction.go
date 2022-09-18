package cmds

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bleuhold/bh/ecsv"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
	"sort"
	"strconv"
	"time"
)

var TRANSACTION *cli.Command
var transactionsFilename = "transactions.json"

var BANKS = map[string]map[int]string{
	"investec": {
		0: "date",
		2: "description",
		3: "debit",
		4: "credit",
		5: "balance",
	},
}

func TransactionsExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case B1:
		ListTransactions()
	}
	return nil
}

// Transaction is a struct containing all fields relevant to a transaction.
//
// Important Note: a transaction is viewed from the perspective from the
// person uploading transactions, that is, for an asset debit increases the
// value and credit decreases the value.
type Transaction struct {
	UUID        uuid.UUID `json:"uuid"`
	AccountUUID uuid.UUID `json:"accountUUID"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	//Items       []Item    `json:"items"`
	Debit      float64   `json:"debit"`
	Credit     float64   `json:"credit"`
	Balance    float64   `json:"balance"`
	Active     bool      `json:"active"`
	CreateDate time.Time `json:"createDate"`
	UpdateDate time.Time `json:"updateDate"`
}

type Transactions []Transaction

// LoadTransactions loads the transactions from the file system.
func LoadTransactions() *Transactions {
	xt := &Transactions{}
	filesys.LoadInterface(transactionsFilename, xt)
	return xt
}

// Save writes the data to the file system.
func (xt *Transactions) Save() {
	sort.Sort(xt)
	xb, err := json.Marshal(xt)
	if err != nil {
		log.Fatalf("unable to marshal transactions data: %v", err)
	}
	filesys.WriteFile(transactionsFilename, xb)
}

// Len returns the length of transactions
func (xt *Transactions) Len() int {
	return len(*xt)
}

// Less returns whether transaction i is before transaction j
func (xt *Transactions) Less(i, j int) bool {
	xti := (*xt)[i]
	xtj := (*xt)[j]
	return xti.Date.Before(xtj.Date)
}

// Swap interchanges two elements in the slice.
func (xt *Transactions) Swap(i, j int) {
	(*xt)[i], (*xt)[j] = (*xt)[j], (*xt)[i]
}

// MarshalCSV reads all the csv rows from the reader and marshals each row into
// its own transaction.
func (xt *Transactions) MarshalCSV(bankCode string, accountUUID uuid.UUID, data *ecsv.CSV) error {
	if bankMap, ok := BANKS[bankCode]; ok {
		for _, rec := range data.Records {
			t := Transaction{
				UUID:        uuid.New(),
				AccountUUID: accountUUID,
				CreateDate:  time.Now(),
				UpdateDate:  time.Now(),
			}
			t.Set(rec, bankMap)
			*xt = append(*xt, t)
		}
		return nil
	}
	return fmt.Errorf("invalid bank code: %v", bankCode)
}

func (xt Transactions) String() string {
	s := fmt.Sprintf("%-36s %-36s %-10s %-40s %-11s %-11s\n", "UUID", "ACCOUNT UUID", "DATE", "DESCRIPTION", "DEBIT", "CREDIT")
	for _, rec := range xt {
		desc := rec.Description
		if len(rec.Description) > 40 {
			desc = rec.Description[:40]
		}
		s += fmt.Sprintf("%-36s %-36s %-10s %-40s %11.2f %11.2f\n", rec.UUID, rec.AccountUUID, rec.Date, desc, rec.Debit, rec.Credit)
	}
	return s
}

// Set is a setter function
func (t *Transaction) Set(record []string, bankMap map[int]string) {
	for key, val := range bankMap {
		v := record[key]
		switch val {
		case "date":
			t.Date, _ = time.Parse("2006/01/02", v)
		case "description":
			t.Description = v
		case "debit":
			t.Debit, _ = strconv.ParseFloat(v, 64)
		case "credit":
			t.Credit, _ = strconv.ParseFloat(v, 64)
		case "balance":
			t.Balance, _ = strconv.ParseFloat(v, 64)
		}
	}
}

// Find finds a transaction based on the UUID. Or returns an error if not found.
// TODO: replace with a hash map for O(1) operation to find transactions.
func (xt *Transactions) Find(UUID uuid.UUID) (*Transaction, error) {
	for _, t := range *xt {
		if t.UUID == UUID {
			return &t, nil
		}
	}
	return &Transaction{}, errors.New("transaction not found")
}

func (xt *Transactions) Remove(UUID uuid.UUID) (*Transactions, error) {
	transactions := *xt
	index := -1
	for i := 0; i < len(transactions); i++ {
		if transactions[i].UUID == UUID {
			index = i
			break
		}
	}
	if index == -1 {
		return xt, fmt.Errorf("transaction not found for uuid: %v", UUID)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)
	return &transactions, nil
}

// ListTransactions loads all the transactions and prints them to the
// transactions.
func ListTransactions() {
	xt := LoadTransactions()
	sort.Sort(xt)
	fmt.Printf("%s", xt.String())
}
