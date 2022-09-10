package cmds

import (
	"fmt"
	"github.com/dottics/cli"
	"github.com/google/uuid"
)

var ACCOUNT_ADD *cli.Command

func AccountAddExecute(cmd *cli.Command) error {
	n, t, p, h, a, err := validateAccount(&S1, &S2, &S3, &S4, &S5)
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case err == nil:
		addAccount(n, t, p, h, a)
	}
	if err != nil {
		cmd.PrintHelp()
		return err
	}
	return nil
}

func validateAccount(number, accountType, providerName, holderName, alias *string) (string, string, string, string, string, error) {
	ok := true
	if *number == "" {
		ok = false
	}
	if *accountType == "" {
		ok = false
	}
	if *providerName == "" {
		ok = false
	}
	if *holderName == "" {
		ok = false
	}
	if *alias == "" {
		ok = false
	}
	if ok {
		return *number, *accountType, *providerName, *holderName, *alias, nil
	} else {
		return *number, *accountType, *providerName, *holderName, *alias, fmt.Errorf("invalid arguments to create a new account")
	}
}

func addAccount(number, accountType, providerName, holderName, alias string) {
	a := NewAccount(number, accountType, providerName, holderName, alias)
	xa := LoadAccounts()
	xa = xa.Add(a)
	xa.Save()
}

var ACCOUNT_REMOVE *cli.Command

func AccountRemoveExecute(cmd *cli.Command) error {
	UUID, err := validateAccountUUID(&S1)
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case err == nil:
		return removeAccount(UUID)
	}
	return err
}

func validateAccountUUID(s *string) (uuid.UUID, error) {
	return uuid.Parse(*s)
}

func removeAccount(UUID uuid.UUID) error {
	xa := LoadAccounts()
	found := false
	removeIndex := -1
	for i, a := range *xa {
		if a.UUID == UUID {
			found = true
			removeIndex = i
		}
	}
	if !found {
		return fmt.Errorf("invalid UUID account not found: %v", UUID)
	} else {
		// remove the account at the remove index
		*xa = append((*xa)[:removeIndex], (*xa)[removeIndex+1:]...)
		xa.Save()
	}
	return nil
}
