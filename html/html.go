package html

import (
	"go/ast"
	"go/parser"
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
			Body: &ast.BlockStmt{
				List: buildInit(infile),
			},
		},
	)
	printer.Fprint(os.Stdout, fset, file)
	return nil
}

func buildInit(infiles []string) []ast.Stmt {
	expr, err := parser.ParseExpr("tmpl.New(\"style.css\").Parse(``)")
	if err != nil {
		panic(err)
	}
	return []ast.Stmt{
		&ast.ExprStmt{
			X: expr,
		},
	}
}
