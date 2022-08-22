package cmds

import (
	"flag"
	"github.com/dottics/cli"
	"log"
)

var INFO *cli.Command
var INFO_SET *cli.Command

func init() {
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
	if err != nil {
		log.Fatalln(err)
	}
}
