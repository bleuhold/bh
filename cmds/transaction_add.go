package cmds

import (
	"errors"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"sort"
	"strings"
	"time"
)

var TRANSACTION_ADD *cli.Command

func TransactionAddExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	default:
		t, err := validateAddTransaction(S1, S2, cmd.FlagSet.Args(), F1, F2)
		if err != nil {
			return err
		}
		err = addTransaction(t)
		return err
	}
}

func validateAddTransaction(S1, S2 string, desc []string, F1, F2 float64) (*Transaction, error) {
	description := strings.Join(desc, " ")
	if description == "" {
		return nil, errors.New("description is a required field")
	}
	accountUUID, err := uuid.Parse(S1)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse("2006-01-02", S2)
	if err != nil {
		return nil, err
	}
	t := &Transaction{
		UUID:        uuid.New(),
		AccountUUID: accountUUID,
		Description: description,
		Date:        date,
		Debit:       F1,
		Credit:      F2,
		Balance:     F2 - F1,
		Active:      true,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
	}
	return t, nil
}

func addTransaction(t *Transaction) error {
	xt := LoadTransactions()
	*xt = append(*xt, *t)
	sort.Sort(xt)

	// now add the item as well
	i := NewItem(t.UUID)
	i.Description = t.Description
	i.Date = t.Date
	i.Debit = t.Debit
	i.Credit = t.Credit
	addItem(i)
	// save transaction
	xt.Save()
	return nil
}
