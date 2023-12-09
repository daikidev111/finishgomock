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

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok || callExpr.Fun == nil {
			return
		}

		// Type assertion for selector expression
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		// Check if the callExpr is to a gomock.Controller method
		if isGomockControllerMethod(pass, callExpr) {
			// Check for the specific Finish method call
			if selectorExpr.Sel.Name == "Finish" {
				// Report
				pass.Reportf(selectorExpr.Sel.NamePos, "detected an unnecessary call to Finish on gomock.Controller")
			}
		}
	})
	return nil, nil
}

// Helper function to check if a call expression is a method on gomock.Controller
func isGomockControllerMethod(pass *analysis.Pass, call *ast.CallExpr) bool {
	if callExpr, ok := call.Fun.(*ast.SelectorExpr); ok {
		receiverType := pass.TypesInfo.TypeOf(callExpr.X)
		if receiverType != nil {
			return strings.HasSuffix(receiverType.String(), "github.com/golang/mock/gomock.Controller")
		}
	}
	return false
}
