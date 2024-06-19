package cmd

import (
	"os"
)

func cmdHelp() {
	if len(os.Args) < 3 {
		usage()
	}
	c, ok := cmds.Find(os.Args[2])
	if ok {
		c.help()
		os.Exit(0)
	}
	usage()
}
