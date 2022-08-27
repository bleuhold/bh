package cmds

import (
	"flag"
	"github.com/dottics/cli"
	"log"
)

var INFO *cli.Command
var INFO_SET *cli.Command

var UPLOAD *cli.Command

var TRANSACTIONS *cli.Command

func init() {
	/*
		INFO
	*/
	INFO = cli.NewCommand("info", &help, flag.ExitOnError)
	INFO.Usage = "bh"
	INFO.Description = "To show the current info."
	INFO.Execute = infoExecute

	INFO.FlagSet.BoolVar(&b1, "list", false, "List all of the current context information.")

	/* SUB COMMANDS */
	INFO_SET = cli.NewCommand("set", &help, flag.ExitOnError)

	INFO_SET.Usage = "bh info"
	INFO_SET.Description = "To set an info parameter."
	INFO_SET.Execute = infoSetExecute

	INFO_SET.FlagSet.StringVar(&s1, "premises-uuid", "", "The premises UUID to be set for filtering.")
	INFO_SET.FlagSet.StringVar(&s2, "start-date", "", "The start date to be set for filtering.")
	INFO_SET.FlagSet.StringVar(&s3, "end-date", "", "The end date to be set for filtering.")

	err := INFO.AddCommands([]*cli.Command{
		INFO_SET,
	})

	/*
		UPLOAD
	*/
	UPLOAD = cli.NewCommand("upload", &help, flag.ExitOnError)
	INFO.Usage = "bh"
	UPLOAD.Description = "Upload a bank statement."
	UPLOAD.Execute = uploadExecute

	UPLOAD.FlagSet.StringVar(&s1, "f", "", "The path to the bank statement CSV to be uploaded.")
	UPLOAD.FlagSet.StringVar(&s1, "file", "", "The path to the bank statement CSV to be uploaded.")

	/*
		TRANSACTIONS
	*/
	TRANSACTIONS = cli.NewCommand("transactions", &help, flag.ExitOnError)
	TRANSACTIONS.Usage = "bh"
	TRANSACTIONS.Description = "All bank transaction related commands."
	TRANSACTIONS.Execute = transactionsExecute

	TRANSACTIONS.FlagSet.BoolVar(&b1, "list", false, "List all transactions for the set date ranges.")

	if err != nil {
		log.Fatalln(err)
	}
}
