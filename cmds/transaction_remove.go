package cmds

import (
	"github.com/dottics/cli"
	"github.com/google/uuid"
)

var TRANSACTION_REMOVE *cli.Command

func TransactionRemoveExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	default:
		UUID, err := uuid.Parse(S1)
		if err != nil {
			return err
		}
		return removeTransaction(UUID)
	}
}

func removeTransaction(UUID uuid.UUID) error {
	xt := LoadTransactions()
	// remove the transaction
	xt, err := xt.Remove(UUID)
	if err != nil {
		return err
	}
	// remove all the transaction items
	xi := LoadItems()
	xi = xi.RemoveTransactionUUID(UUID)
	// save both the items and the transactions
	xi.Save()
	xt.Save()
	return nil
}
