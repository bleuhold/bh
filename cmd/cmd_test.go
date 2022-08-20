package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestCommand_Help(t *testing.T) {
	c := Command{
		Name:        "get",
		Description: "Get all <entries>",
	}

	expected := fmt.Sprintf("Usage: bh %s\n\n%s\n\n", c.Name, c.Description)
	helpString := c.Help()
	if helpString != expected {
		t.Errorf("expected help string '%v' got '%v'", expected, helpString)
	}
}

func TestCommands_Help(t *testing.T) {
	xc := Commands{
		"get": &Command{
			Name:        "get",
			Description: "Get all <entries>",
		},
		"adds": &Command{
			Name:        "adds",
			Description: "Add a new <entry>",
		},
	}

	expected := fmt.Sprintf("Usage: bh COMMAND\n\nCommands:\n  get         Get all <entries>\n  adds        Add a new <entry>\n")
	helpString := xc.Help()
	if helpString != expected {
		t.Errorf("expected help string '%v' got '%v'", expected, helpString)
	}
}

func TestCommandSet_Add(t *testing.T) {
	cs := CommandSet{
		Commands: map[string]*Command{
			"add": {
				Name: "add",
			},
		},
	}

	tt := []struct {
		name  string
		cmd   Command
		error error
	}{
		{
			name: "fail as command already exists",
			cmd: Command{
				Name: "add",
			},
			error: errors.New("cannot add command add already exists"),
		},
		{
			name: "add new command",
			cmd: Command{
				Name: "remove",
			},
			error: nil,
		},
	}

	// this test simply ensures that when a command is added that it
	// returns an error, when the same command is added.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.Add(&tc.cmd)
			err1 := fmt.Sprintf("%v", err)
			err2 := fmt.Sprintf("%v", tc.error)
			if err1 != err2 {
				t.Errorf("expected error '%v' got '%v'", err2, err1)
			}
		})
	}
}

func TestCommandSet_Run_Level_0(t *testing.T) {
	// best practise to restore the global state to as before
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cs := &CommandSet{
		level: 0,
		Commands: map[string]*Command{
			"add": {
				Name:    "add",
				FlagSet: flag.NewFlagSet("add", flag.ExitOnError),
			},
		},
	}

	os.Args = []string{"bh", "add", "-help"}
	fmt.Printf("cs: %v\n", cs)

	c, err := cs.Run()
	fmt.Printf("%v\n", c)
	if err != nil {
		t.Errorf("expected error nil got %v", err)
	}
	cmd := cs.Commands["add"]
	if c == nil {
		t.Errorf("expected command got nil")
	} else {
		if c != cmd {
			t.Errorf("expected command address %v got %v", cmd, c)
		}
	}
}
