package main

import (
	"go/ast"
	"go/parser"
)

func main() {
	expr, _ := parser.ParseExpr("A + 1")

	ast.Print(nil, expr)
}
