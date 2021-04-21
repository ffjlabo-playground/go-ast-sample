package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/helloworld.pb.go", nil, parser.Mode(0))

	structASTMap := map[string]*ast.StructType{}

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

			structASTMap[s.Name.Name] = t
		}
	}

	// structの各要素を取得する
	for name, node := range structASTMap {
		// struct名
		fmt.Println(name)
		for _, field := range node.Fields.List {

			for _, nameIdent := range field.Names {
				// 要素名
				fmt.Print("    " + nameIdent.Name + " ")

				// 要素の型名
				switch field.Type.(type) {
				// 別パッケージの型を利用している場合
				case *ast.SelectorExpr:
					selectorExpr, _ := field.Type.(*ast.SelectorExpr)
					xIdent, _ := selectorExpr.X.(*ast.Ident)
					fmt.Println(xIdent.Name + "." + selectorExpr.Sel.Name)
				// 組み込みまたはどうパッケージ内の型を利用している場合
				case *ast.Ident:
					ident, _ := field.Type.(*ast.Ident)
					fmt.Println(ident.Name)
				}
				// typeSelectorExpr, _ := field.Type.(*ast.SelectorExpr)
				// typeIdent, _ := typeSelectorExpr.X.(*ast.Ident)
				// fmt.Println(field.Type)
			}
		}
	}
}
