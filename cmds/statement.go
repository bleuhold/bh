package cmds

import (
	"fmt"
	"github.com/dottics/cli"
	"github.com/google/uuid"
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
	case S1 != "":
		UUID, err := uuid.Parse(S1)
		if err != nil {
			return err
		}
		c := GetContract(UUID)
		xi := LoadItems()
		items := xi.DateRange(c.Dates.Occupation, c.Dates.Termination)
		items = items.FilterTags(c.References())

		fmt.Printf("\n\u001B[1mSTATEMENT\u001B[0m\n\n")
		c.Print(false)
		fmt.Println(items.String())
	}
	return nil
}
