package main

import (
	"fmt"
	"github.com/bleuhold/bh/cmd"
	"github.com/bleuhold/bh/cmds"
	"github.com/bleuhold/bh/filesys"
	"log"
	"os"
)

func main() {
	// set up the default data directory for the command line tool.
	filesys.DirectorySetup()
	fmt.Printf("**%v**\n", os.Args)
	cs := cmd.NewCommandSet()

	err := cs.Add(cmds.INFO)
	if err != nil {
		log.Fatalln(err)
	}
	err = cs.Add(cmds.PREMISES)
	if err != nil {
		log.Fatalln(err)
	}

	err = cs.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

//
//func main() {
//	//fmt.Println("Hello")
//	// if there is an error the "get" subcommand will exit on error
//	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
//	var getAll bool
//	getCmd.BoolVar(&getAll, "all", false, "")
//	//getAll := getCmd.Bool("all", false, "Get all videos.")
//	var getID string
//	getCmd.StringVar(&getID, "id", "", "YouTube video ID.")
//	getCmd.StringVar(&getID, "I", "", "YouTube video ID.")
//
//	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
//	addID := addCmd.String("id", "", "YouTube video ID.")
//	addTitle := addCmd.String("title", "", "YouTube video title.")
//
//	if len(os.Args) < 2 {
//		fmt.Println("expected 'get' or 'add' subcommand")
//		os.Exit(1)
//	}
//
//	switch os.Args[1] {
//	case "get":
//		// handle get
//		HandleGet(getCmd, &getAll, &getID)
//	case "add":
//		// handle add
//		HandleAdd(addCmd, addID, addTitle)
//	default:
//		// handle default
//	}
//}
//
//func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
//	err := getCmd.Parse(os.Args[2:])
//	if err != nil {
//		log.Fatalln("Unable to parse get command", err)
//	}
//
//	if *all == false && *id == "" {
//		fmt.Println("id is required or specify --all for all videos")
//		getCmd.PrintDefaults()
//		os.Exit(1)
//	}
//
//	fmt.Println("get all videos")
//}
//
//func HandleAdd(addCmd *flag.FlagSet, id *string, title *string) {}
