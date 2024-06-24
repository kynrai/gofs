package cmd

import (
	"fmt"
	folder "module/placeholder"
	"os"

	"github.com/atos-digital/ttz/internal/gen"
)

const (
	root              = "template"
	defaultModuleName = "module/placeholder"
)

const initUsage = `usage: 10100 init [module-name] [dir]

"init" initializes a new module in the specified directory.
If no directory is specified, the current directory is used.

The module name should be a go module name, e.g. "github.com/user/module".

Example:
  10100 init mymodule /path/to/dir
  10100 init mymodule

`

func init() {
	Ttz.AddCmd(Command{
		Name:  "init",
		Short: "initialize a new 10100 module",
		Long:  initUsage,
		Cmd:   cmdInit,
	})
}

func cmdInit() {
	args := os.Args[2:] // skip program name and command

	moduleName := ""
	dir := ""
	var err error

	switch {
	case len(args) == 0:
		fmt.Println("init: missing module name")
		fmt.Print(initUsage)
		return
	case len(args) == 1:
		moduleName = args[0]
		dir, err = os.Getwd()
		if err != nil {
			fmt.Println("init: ", err)
			return
		}
	case len(args) == 2:
		moduleName = args[0]
		dir = args[1]
	default:
		fmt.Println("init: too many arguments")
		fmt.Print(initUsage)
		return
	}

	parser := gen.NewParser(dir, defaultModuleName, moduleName, folder.Folder)
	parser.Parse()
}
