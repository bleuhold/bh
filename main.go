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

func main() {
	// set up the default data directory for the command line tool.
	filesys.DirectorySetup()
	fmt.Printf("**%v**\n\n", os.Args)

	c := cli.NewCommand("bh", flag.ExitOnError)
	c.Execute = executeBH

	// add commands
	err := c.AddCommands([]*cli.Command{
		cmds.INFO,
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}

func executeBH(c *cli.Command) error {
	c.PrintHelp()
	return nil
}
