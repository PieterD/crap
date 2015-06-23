package html

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PieterD/crap/agoge/internal/build"
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
	var list []ast.Stmt

	for _, filename := range infiles {
		bstr, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		str := strings.Replace(string(bstr), "`", "`+\"`\"+`", -1)
		expr, err := parser.ParseExpr("tmpl.New(\"style.css\").Parse(`" + str + "`)")
		if err != nil {
			panic(err)
		}
		list = append(list, &ast.ExprStmt{
			X: expr,
		})
	}
	return list
}
