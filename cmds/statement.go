package cmds

import (
	"github.com/dottics/cli"
)

var STATEMENT *cli.Command

/*
	STATEMENT
*/

func StatementExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	}
	return nil
}
