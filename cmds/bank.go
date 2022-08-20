package cmds

import (
	"flag"
	"github.com/bleuhold/bh/cmd"
)

var BANK *cmd.Command

//var bankFilename = "bank.json"

func init() {
	BANK = cmd.NewCommand("bank", flag.ExitOnError)
}
