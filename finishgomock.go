package finishgomock

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "finishgomock is a linter that detects an unnecessary call to Finish on gomock.Controller"

var Analyzer = &analysis.Analyzer{
	Name: "finishgomock",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	var flgMockPkg bool // once true, it stays true for the rest of the preorder traversal process

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch callExprTyp := n.(type) {
		case *ast.CallExpr:
			if strings.Contains(pass.TypesInfo.TypeOf(callExprTyp).String(), "github.com/golang/mock/gomock.Controller") {
				flgMockPkg = true
			}
			if callExprTyp.Fun == nil {
				return
			}
			fieldSel := callExprTyp.Fun.(*ast.SelectorExpr).Sel
			if fieldSel.Name == "Finish" && flgMockPkg {
				pass.Reportf(fieldSel.NamePos, "detected an unnecessary call to Finish on gomock.Controllers")
			}
		}
	})
	return nil, nil
}
