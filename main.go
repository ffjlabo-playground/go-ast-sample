package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	for _, decl := range f.Decls {
		// struct型の型定義を表示する

		// GenDeclであるかを判定
		d, ok := decl.(*ast.GenDecl)
		if !ok || d.Tok != token.TYPE {
			continue
		}

		for _, spec := range d.Specs {
			// TypeSpec型か確認
			s, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// StructType型か確認
			t, ok := s.Type.(*ast.StructType)
			if !ok {
				continue
			}
			fmt.Println(s.Name)
			ast.Print(fset, t)
		}
	}
}
