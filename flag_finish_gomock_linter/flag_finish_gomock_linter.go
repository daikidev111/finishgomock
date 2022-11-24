package flag_finish_gomock_linter

import (
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "flag_finish_gomock_linter is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "flag_finish_gomock_linter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
    fset := token.NewFileSet()
    f, err := parser.ParseFile(fset, "./testdata/src/a/a.go", nil, 0)
	if err != nil {
		panic(err)
	}
    ast.Print(fset, f)

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "gopher" {
				pass.Reportf(n.Pos(), "identifier is gopher")
			}
		}
	})

	return nil, nil
}
