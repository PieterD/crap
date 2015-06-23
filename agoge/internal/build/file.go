package build

import "go/ast"

func NewFile(pkg string, decl ...ast.Decl) *ast.File {
	return &ast.File{
		Name: &ast.Ident{
			Name: pkg,
		},
		Decls: decl,
	}
}
