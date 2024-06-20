package main

import (
	"github.com/atos-digital/10100-cli/internal/cli"
	"github.com/atos-digital/10100-cli/internal/gen"
)

const (
	root              = "template"
	defaultModuleName = "module/placeholder2"
)

func main() {
	cli.New()
	parser := gen.NewParser(root, defaultModuleName, "module/placeholder")
	parser.Parse()
}
