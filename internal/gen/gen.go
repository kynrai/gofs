package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"golang.org/x/tools/go/ast/astutil"
)

type Parser struct {
	// DirPath is the path to the folder to parse.
	// This should be a directory containing the go files to parse.
	DirPath string
	Files   map[string]*ast.Package
}

func NewParser(dirPath string) *Parser {
	return &Parser{DirPath: dirPath}
}

func (p *Parser) Parse() error {
	return nil
}

func (p *Parser) renameModule(name string) error {
	return nil
}

func (p *Parser) parseDir() error {
	return nil
}

func (p *Parser) parseFile() error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "template/main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.CallExpr:
			for i, arg := range x.Args {
				a, ok := arg.(*ast.BasicLit)
				x.Args[i] = &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"new argument"`,
				}
				if ok {
					fmt.Println(a.Value)
				}
			}
			id, ok := x.Fun.(*ast.SelectorExpr)
			if ok {
				fmt.Println(id.Sel.Name, id.Sel.NamePos)
			}
		}

		return true
	})
	return nil
}
