package cmds

import "github.com/dottics/cli"

// uploadExecute is the function executed when the upload command is called.
func uploadExecute(cmd *cli.Command) error {
	switch {
	case help:
		cmd.PrintHelp()
		return nil
	}
	return nil
}
