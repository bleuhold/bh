package cmd

import (
	"errors"
	"fmt"
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
		"get": Command{
			Name:        "get",
			Description: "Get all <entries>",
		},
		"adds": Command{
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
		Commands: map[string]Command{
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
