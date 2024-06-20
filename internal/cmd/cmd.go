package cmd

import (
	"fmt"
	"os"
	"sort"
)

type command struct {
	name        string
	description string
	help        func()
	fn          func()
}

type commands []command

func (c commands) Find(cmd string) (command, bool) {
	for _, v := range c {
		if v.name == cmd {
			return v, true
		}
	}
	return command{}, false
}

func (c *commands) AddCmd(cmd command) {
	*c = append(*c, cmd)
}

var cmds = commands{}

func New() {
	if len(os.Args) < 2 {
		usage()
	}
	cmd := os.Args[1]

	// special case for help
	if cmd == "help" {
		cmdHelp()
		os.Exit(0)
	}

	c, ok := cmds.Find(cmd)
	if !ok {
		usage()
	}
	c.fn()
	os.Exit(0)
}

// usage returns the usage message, always sorted by command name
// always exits the program with status code 0
func usage() {
	msg := `10100 is a tool for managing 10100 projects.
	
Usage:

	10100 <command> [arguments]

The commands are:

%s
Use "10100 help <command>" for more information about a command.

`
	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].name < cmds[j].name
	})
	cmdList := ""
	for _, c := range cmds {
		cmdList += fmt.Sprintf("\t%s\t\t%s\n", c.name, c.description)
	}
	fmt.Printf(msg, cmdList)
	os.Exit(0)
}
