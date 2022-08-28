package cmds

import (
	"fmt"
	"github.com/dottics/cli"
)

var ACCOUNT_ADD *cli.Command

func AccountAddExecute(cmd *cli.Command) error {
	n, t, p, err := validateAccount(&S1, &S2, &S3)
	switch {
	case Help:
		cmd.PrintHelp()
	case err == nil:
		addAccount(n, t, p)
	}
	if err != nil {
		cmd.PrintHelp()
		return err
	}
	return nil
}

func validateAccount(number, accountType, providerName *string) (string, string, string, error) {
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
	if ok {
		return *number, *accountType, *providerName, nil
	} else {
		return *number, *accountType, *providerName, fmt.Errorf("invalid arguments to create a new account")
	}
}

func addAccount(number, accountType, providerName string) {
	a := NewAccount(number, accountType, providerName)
	xa := LoadAccounts()
	xa = xa.Add(a)
	xa.Save()
}
