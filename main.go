package main

import (
	"flag"
	"fmt"
	"github.com/bleuhold/bh/cmds"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"log"
	"os"
)

var help bool

//Black: \u001b[30m
//Red: \u001b[31m
//Green: \u001b[32m
//Yellow: \u001b[33m
//Blue: \u001b[34m
//Magenta: \u001b[35m
//Cyan: \u001b[36m
//White: \u001b[37m
//Reset: \u001b[0m

func init() {
	/*
		INFO
	*/
	cmds.INFO = cli.NewCommand("info", &cmds.Help, flag.ExitOnError)
	cmds.INFO.Usage = "bh"
	cmds.INFO.Description = "To show the current info."
	cmds.INFO.Execute = cmds.InfoExecute

	cmds.INFO.FlagSet.BoolVar(&cmds.B1, "list", false, "List all of the current context information.")

	/* SUB COMMANDS */
	cmds.INFO_SET = cli.NewCommand("set", &cmds.Help, flag.ExitOnError)

	cmds.INFO_SET.Usage = "bh info"
	cmds.INFO_SET.Description = "To set an info parameter."
	cmds.INFO_SET.Execute = cmds.InfoSetExecute

	cmds.INFO_SET.FlagSet.StringVar(&cmds.S1, "premises-uuid", "", "The premises UUID to be set for filtering.")
	cmds.INFO_SET.FlagSet.StringVar(&cmds.S2, "start-date", "", "The start date to be set for filtering.")
	cmds.INFO_SET.FlagSet.StringVar(&cmds.S3, "end-date", "", "The end date to be set for filtering.")

	err := cmds.INFO.AddCommands([]*cli.Command{
		cmds.INFO_SET,
	})

	/*
		UPLOAD
	*/
	cmds.UPLOAD = cli.NewCommand("upload", &cmds.Help, flag.ExitOnError)
	cmds.UPLOAD.Usage = "bh"
	cmds.UPLOAD.Description = "Upload a bank statement."
	cmds.UPLOAD.Execute = cmds.UploadExecute

	cmds.UPLOAD.FlagSet.StringVar(&cmds.S1, "f", "", "The path to the bank statement CSV to be uploaded.")
	cmds.UPLOAD.FlagSet.StringVar(&cmds.S1, "file", "", "The path to the bank statement CSV to be uploaded.")
	cmds.UPLOAD.FlagSet.StringVar(&cmds.S2, "uuid", "", "The wallet/account UUID for the transactions to be associated with.")

	/*
		TRANSACTIONS
	*/
	cmds.TRANSACTIONS = cli.NewCommand("transactions", &cmds.Help, flag.ExitOnError)
	cmds.TRANSACTIONS.Usage = "bh"
	cmds.TRANSACTIONS.Description = "All bank transaction related commands."
	cmds.TRANSACTIONS.Execute = cmds.TransactionsExecute

	cmds.TRANSACTIONS.FlagSet.BoolVar(&cmds.B1, "list", false, "List all transactions for the set date ranges.")

	/*
		ACCOUNT
	*/
	cmds.ACCOUNT = cli.NewCommand("account", &cmds.Help, flag.ExitOnError)
	cmds.ACCOUNT.Usage = "bh"
	cmds.ACCOUNT.Description = "All wallet/bank accounts."
	cmds.ACCOUNT.Execute = cmds.AccountExecute

	cmds.ACCOUNT.FlagSet.BoolVar(&cmds.B1, "list", false, "List all accounts.")

	cmds.ACCOUNT_ADD = cli.NewCommand("add", &cmds.Help, flag.ExitOnError)
	cmds.ACCOUNT_ADD.Usage = "bh account"
	cmds.ACCOUNT_ADD.Description = "Add a new account."
	cmds.ACCOUNT_ADD.Execute = cmds.AccountAddExecute

	cmds.ACCOUNT_ADD.FlagSet.StringVar(&cmds.S1, "number", "", "The wallet/account number or identifier.")
	cmds.ACCOUNT_ADD.FlagSet.StringVar(&cmds.S2, "type", "", "The wallet/account type.")
	cmds.ACCOUNT_ADD.FlagSet.StringVar(&cmds.S3, "provider", "", "The wallet/account provider (bank/organisation).")

	cmds.ACCOUNT_REMOVE = cli.NewCommand("remove", &cmds.Help, flag.ExitOnError)
	cmds.ACCOUNT_REMOVE.Usage = "bh account"
	cmds.ACCOUNT_REMOVE.Description = "Remove an account."
	cmds.ACCOUNT_REMOVE.Execute = cmds.AccountRemoveExecute

	cmds.ACCOUNT_REMOVE.FlagSet.StringVar(&cmds.S1, "uuid", "", "The wallet/account uuid to be removed.")

	err = cmds.ACCOUNT.AddCommands([]*cli.Command{
		cmds.ACCOUNT_ADD,
		cmds.ACCOUNT_REMOVE,
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// set up the default data directory for the command line tool.
	filesys.DirectorySetup()
	fmt.Printf("**%v**\n\n", os.Args)

	c := cli.NewCommand("bh", &help, flag.ExitOnError)
	c.Execute = executeBH

	// add commands
	err := c.AddCommands([]*cli.Command{
		cmds.INFO,
		cmds.UPLOAD,
		cmds.TRANSACTIONS,
		cmds.ACCOUNT,
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Run(os.Args[1:])
	if err != nil {
		fmt.Printf("\u001b[31mERROR\u001b[0m: %v\n\n", err)
	}
}

func executeBH(c *cli.Command) error {
	c.PrintHelp()
	return nil
}
