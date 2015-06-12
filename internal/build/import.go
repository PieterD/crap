package build

import (
	"go/ast"
	"go/token"
)

func ImportList(imports ...*ast.ImportSpec) *ast.GenDecl {
	var list []ast.Spec
	for _, imp := range imports {
		list = append(list, imp)
	}
	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: list,
	}
}

func Import(name, path string) *ast.ImportSpec {
	var ident *ast.Ident
	if name != "" {
		ident = &ast.Ident{
			Name: name,
		}
	}
	return &ast.ImportSpec{
		Name: ident,
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"" + path + "\"",
		},
	}
}
