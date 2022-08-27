package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/ecsv"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

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

func transactionsExecute(cmd *cli.Command) error {
	switch {
	case help:
		cmd.PrintHelp()
		return nil
	case b1:
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
	xb, err := json.Marshal(xt)
	if err != nil {
		log.Fatalf("unable to marshal transactions data: %v", err)
	}
	filesys.WriteFile(transactionsFilename, xb)
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
	s := ""
	for _, rec := range xt {
		desc := rec.Description
		if len(rec.Description) > 40 {
			desc = rec.Description[:41]
		}
		s += fmt.Sprintf("%36s  ", rec.AccountUUID)
		s += fmt.Sprintf("%1s  ", rec.Date.Format("2006-01-02"))
		s += fmt.Sprintf("%-40s  ", desc)
		s += fmt.Sprintf("%11.2f  ", rec.Debit)
		s += fmt.Sprintf("%11.2f  ", rec.Credit)
		s += fmt.Sprintf("%11.2f", rec.Balance)
		s += "\n"
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

// ListTransactions loads all the transactions and prints them to the
// transactions.
func ListTransactions() {
	xt := LoadTransactions()
	fmt.Printf("%s", xt.String())
}
