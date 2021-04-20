package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			// ノードを書き換える
			v.Name = &ast.Ident{
				Name: "plus",
			}
		}

		return true
	})

	file, err := os.OpenFile("example/result.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(file, fset, f)
}
