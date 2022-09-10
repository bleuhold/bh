package cmds

import (
	"fmt"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

var ITEM_ADD *cli.Command

// ItemAddExecute is the function executed when the command `bh item add` is
// run. To add a new item to a transaction.
func ItemAddExecute(cmd *cli.Command) error {
	item, err := validateItemAdd(S1, S2, S3, cmd.FlagSet.Args())
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case err == nil:
		addItem(item)
	}
	return nil
}

func validateItemAdd(s1, s2, s3 string, desc []string) (*Item, error) {
	description := strings.Join(desc, " ")
	if description == "" {
		return &Item{}, fmt.Errorf("description is a required field")
	}
	UUID, err := uuid.Parse(s1)
	if err != nil {
		return &Item{}, err
	}
	xt := LoadTransactions()
	t, err := xt.Find(UUID)
	if err != nil {
		return &Item{}, err
	}
	debits, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return &Item{}, err
	}
	credits, err := strconv.ParseFloat(s3, 64)
	if err != nil {
		return &Item{}, err
	}
	i := NewItem(t.UUID)
	i.Description = description
	i.Date = t.Date
	i.Debit = debits
	i.Credit = credits
	return i, nil
}

func addItem(i *Item) {
	xi := LoadItems()
	xi.Add(*i)
	xi.Save()
}
