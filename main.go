package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"

	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "master/main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.CallExpr:
			fmt.Println(reflect.TypeOf(x.Args))
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

	fmt.Println("Modified AST:")
	format.Node(os.Stdout, fset, file)

}
