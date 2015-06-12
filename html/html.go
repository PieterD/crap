package html

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"github.com/PieterD/agoge/internal/build"
)

func Incorporate(pkg string, outfile string, infile ...string) error {
	fset := token.NewFileSet()
	file := build.NewFile(
		pkg,
		build.ImportList(
			build.Import("", "html/template"),
		),
		&ast.FuncDecl{
			Name: &ast.Ident{
				Name: "init",
			},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{},
				},
			},
			//Body: &ast.BlockStmt{
			//	List: []ast.Stmt{},
			//},
		},
	)
	printer.Fprint(os.Stdout, fset, file)
	return nil
}
