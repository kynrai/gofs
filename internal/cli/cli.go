package cli

import (
	"flag"
	"fmt"
)

type Cli struct {
	name     string
	commands []Command
}

// New initiates the CLI command handlers that handle args and params
func New() {
	c := Cli{
		name: "10100-cli",
		commands: []Command{
			&helpCommand{},
		},
	}
	c.cmd()
}

func (c Cli) cmd() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		c.help()
		return
	}
}

func (c Cli) help() {
	fmt.Printf("Usage: %s [command]\n\n", c.name)
	for _, cmd := range c.commands {
		fmt.Printf("\t%s\t%s\n", cmd.Name(), cmd.Description())
	}
	fmt.Println()
}
