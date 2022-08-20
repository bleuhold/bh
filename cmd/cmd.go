package cmd

import (
	"flag"
	"fmt"
	"os"
)

func WIP(cmd *Command) {
	fmt.Printf("WIP\n\n")
}

// Command is a struct
type Command struct {
	Name        string
	Description string
	FlagSet     *flag.FlagSet
	//CommandSet  *CommandSet
	Execute func(command *Command)
}

// NewCommand creates a basic new command.
func NewCommand(name string, handling flag.ErrorHandling) *Command {
	cmd := &Command{
		Name:        name,
		Description: "",
		FlagSet:     flag.NewFlagSet(name, handling),
		Execute:     WIP,
	}
	return cmd
}

// Help is the method that prints the help description of the command to
// the Standard Output.
func (c *Command) Help() string {
	return fmt.Sprintf("Usage: bh %s\n\n%s\n\n", c.Name, c.Description)
}

// Init parses the command line args to the command's flags.
func (c *Command) Init(args []string) error {
	return c.FlagSet.Parse(args)
}

// PrintHelp prints the command help to the console.
func (c Command) PrintHelp() {
	fmt.Printf("%s", c.Help())
	c.FlagSet.PrintDefaults()
}

// Commands is a slice of all commands for the command line tool
type Commands map[string]*Command

// Help formats the string that is printed to the os.StdOut when
// the --help flag is passed.
func (c Commands) Help() string {
	s := "Usage: bh COMMAND\n\nCommands:\n"
	for _, cmd := range c {
		s += fmt.Sprintf("  %-10s  %s\n", cmd.Name, cmd.Description)
	}
	return s
}

// PrintHelp prints the commands help to the console.
func (c Commands) PrintHelp() {
	fmt.Println(c.Help())
}

type CommandSet struct {
	level    int
	Commands Commands
}

// NewCommandSet creates a new CommandSet.
func NewCommandSet(level int) *CommandSet {
	cs := &CommandSet{
		level: level,
	}
	cs.Commands = make(map[string]*Command)
	return cs
}

// Add appends a command to the command set.
func (cs *CommandSet) Add(cmd *Command) error {
	// Fatal if the command already exists
	if _, ok := cs.Commands[cmd.Name]; ok {
		return fmt.Errorf("cannot add command %s already exists", cmd.Name)
	}
	cs.Commands[cmd.Name] = cmd
	return nil
}

// Run handles the execution of the command set.
func (cs *CommandSet) Run() (*Command, error) {
	fmt.Printf("args: %v\n", os.Args)
	//help := flag.Bool("help", false, "Show help")
	//if len(os.Args) < 2 {
	//	cs.Commands.PrintHelp()
	//	return nil, nil
	//}
	//flag.Parse()
	if len(os.Args) < 2 {
		cs.Commands.PrintHelp()
		return nil, nil
	}

	command := os.Args[1+cs.level]
	args := os.Args[(2 + cs.level):]
	fmt.Printf("%v %v\n", command, args)

	cmd, ok := cs.Commands[command]
	switch {
	//case *help:
	//	cs.Commands.PrintHelp()
	case ok:
		if err := cmd.Init(args); err != nil {
			return nil, err
		}
		//cmd.Execute(&cmd)
		return cmd, nil
	default:
		fmt.Printf("Invalid command: %s\n\n", os.Args[1])
		cs.Commands.PrintHelp()
	}
	return nil, nil
}
