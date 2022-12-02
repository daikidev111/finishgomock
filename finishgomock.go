package finishgomock

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "finishgomock is a linter that detects a GoMock Finish call when the testing package is used"

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
			var flgMockFinish bool
			if strings.Contains(pass.TypesInfo.TypeOf(callExprTyp).String(), "github.com/golang/mock/gomock.Controller") {
				flgMockPkg = true
			}
			if callExprTyp.Fun == nil {
				return
			}
			selExpr, ok := callExprTyp.Fun.(*ast.SelectorExpr)
			if !ok { // if target node is not detected, break to prevent from throwing panic
				break
			}
			if selExpr.Sel.Name == "Finish" {
				flgMockFinish = true
			}
			if flgMockFinish && flgMockPkg { // if both true, finish and gomock used
				pass.Reportf(selExpr.Sel.NamePos, "identifier is GoMock Finish")
			}
		}
	})
	return nil, nil
}
