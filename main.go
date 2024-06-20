package main

import (
	"github.com/atos-digital/10100-cli/internal/cmd"
)

const (
	root              = "template"
	defaultModuleName = "module/placeholder"
)

func main() {
	cmd.New()
	// parser := gen.NewParser(root, defaultModuleName, "module/placeholder2")
	// parser.Parse()
}
