package finishgomock

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "finishgomock is a linter that detects a GoMock Finish call when the testing package is used"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "finishgomock",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	var flgMockPkg bool // once true, it stays true for the rest of preorder traversal process

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch callExprTyp := n.(type) {
		case *ast.CallExpr:
			var flgMockFinish bool

			// new_cont_flag := false
			if (strings.Contains(pass.TypesInfo.TypeOf(callExprTyp).String(), "github.com/golang/mock/gomock.Controller")) {
				flgMockPkg = true
				// if (strings.Compare(callExprTyp.Fun.(*ast.SelectorExpr).Sel.Name, "NewController") == 0) {
				// 	new_cont_flag = true
				// }
			}

			if callExprTyp.Fun == nil {
				return
			}
			selExpr, ok := callExprTyp.Fun.(*ast.SelectorExpr)
			if !ok { // if target node is not detected, break to prevent from throwing panic
				break
			}
			if (selExpr.Sel.Name == "Finish") {
				flgMockFinish = true
			}

			if (flgMockFinish && flgMockPkg) { // if both true, finish and gomock used
				pass.Reportf(selExpr.Sel.NamePos, "identifier is GoMock Finish") 
			}

		} 
	})

	return nil, nil
}


// func run(pass *analysis.Pass) (any, error) {
//     // fset := token.NewFileSet()
//     // f, err := parser.ParseFile(fset, "./testdata/src/a/a.go", nil, 0)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
//     // ast.Print(fset, f)

// 	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

// 	nodeFilter := []ast.Node{
// 		(*ast.FuncDecl)(nil),
// 	}

// 	inspect.Preorder(nodeFilter, func(n ast.Node) {
// 		switch n := n.(type) {
// 		case *ast.FuncDecl:
// 				if n.Type.Params.List == nil {
// 					return
// 				} 
// 				for _, l := range n.Type.Params.List {
// 					starExp, ok := l.Type.(*ast.StarExpr)
// 					if !ok {
// 						continue
// 					}

// 					arg_test := false
// 					if starExp.X.(*ast.SelectorExpr).X.(*ast.Ident).Name == "testing" { arg_test = true }

// 					if n.Body.List == nil {
// 						return
// 					}

// 					for k, i := range n.Body.List { // can optimise?
// 						// TODO: Fix if it is a different ast node type otherwise it would throw panic error 
// 						// TODO: Remove k for indexing
// 						if k == 1 {
// 							if (i.(*ast.DeferStmt).Call.Fun.(*ast.SelectorExpr).Sel.Name == "Finish" && arg_test) {
// 								pass.Reportf(i.(*ast.DeferStmt).Call.Fun.(*ast.SelectorExpr).Sel.NamePos, "identifier is GoMock Finish")
// 							}
// 							fmt.Println(reflect.TypeOf(i.(*ast.DeferStmt).Call.Args))
// 						}
// 						// log.Print(k, i)
// 					}
// 				}

// 		}
// 	})

// 	return nil, nil
// }
