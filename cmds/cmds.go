package cmds

import (
	"flag"
	"github.com/dottics/cli"
	"log"
)

var INFO *cli.Command
var INFO_SET *cli.Command

func init() {
	INFO = cli.NewCommand("info", flag.ExitOnError)
	INFO.Usage = "bh"
	INFO.Description = "To show the current info."
	INFO.Execute = infoExecute

	INFO.FlagSet.BoolVar(&list, "list", false, "List all of the current context information.")

	/* SUB COMMANDS */
	INFO_SET = cli.NewCommand("set", flag.ExitOnError)

	INFO_SET.Usage = "bh info"
	INFO_SET.Description = "To set an info parameter."
	INFO_SET.Execute = infoSetExecute

	INFO_SET.FlagSet.StringVar(&premisesUUID, "premises-uuid", "", "The premises UUID to be set for filtering.")
	INFO_SET.FlagSet.StringVar(&startDate, "start-date", "", "The start date to be set for filtering.")
	INFO_SET.FlagSet.StringVar(&endDate, "end-date", "", "The end date to be set for filtering.")

	err := INFO.AddCommands([]*cli.Command{
		INFO_SET,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
