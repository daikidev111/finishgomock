package flag_finish_gomock_linter

import (
	"fmt"
	"go/ast"
	"reflect"

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

// func run(pass *analysis.Pass) (any, error) {
//     // fset := token.NewFileSet()
//     // f, err := parser.ParseFile(fset, "./testdata/src/a/a.go", nil, 0)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
//     // ast.Print(fset, f)

// 	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

// 	nodeFilter := []ast.Node{
// 		(*ast.CallExpr)(nil),
// 		(*ast.ExprStmt)(nil), 
// 		(*ast.SelectorExpr)(nil), 
// 	}

// 	inspect.Preorder(nodeFilter, func(n ast.Node) {
// 		new_cont_flag := false
// 		finish_flag := false
// 		switch callExprTyp := n.(type) {
// 		case *ast.CallExpr:
// 			if (strings.Contains(pass.TypesInfo.TypeOf(callExprTyp).String(), "github.com/golang/mock/gomock.Controller")) {
// 				// ast.Print(fset, callExprTyp)
// 				if callExprTyp.Fun == nil {
// 					return 
// 				}
// 				if (strings.Compare(callExprTyp.Fun.(*ast.SelectorExpr).Sel.Name, "NewController") == 0) {
// 					new_cont_flag = true
// 				}
// 				// print(callExprTyp.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Name) // gomock
// 			}
// 			// ast.Print(fset, callExprTyp)
// 		case *ast.SelectorExpr:
// 			if strings.Compare(callExprTyp.Sel.Name, "Finish") == 0 {
// 				finish_flag = true
// 			}
// 			log.Print(finish_flag, new_cont_flag)
// 		} 
// 	})

// 	return nil, nil
// }


func run(pass *analysis.Pass) (any, error) {
    // fset := token.NewFileSet()
    // f, err := parser.ParseFile(fset, "./testdata/src/a/a.go", nil, 0)
	// if err != nil {
	// 	panic(err)
	// }
    // ast.Print(fset, f)

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
				if n.Type.Params.List == nil {
					return
				} 
				for _, l := range n.Type.Params.List {
					starExp, ok := l.Type.(*ast.StarExpr)
					if !ok {
						continue
					}

					arg_test := false
					if starExp.X.(*ast.SelectorExpr).X.(*ast.Ident).Name == "testing" { arg_test = true }

					if n.Body.List == nil {
						return
					}

					for k, i := range n.Body.List { // can optimise?
						// TODO: Fix if it is a different ast node type otherwise it would throw panic error 
						// TODO: Remove k for indexing
						if k == 1 {
							if (i.(*ast.DeferStmt).Call.Fun.(*ast.SelectorExpr).Sel.Name == "Finish" && arg_test) {
								pass.Reportf(i.(*ast.DeferStmt).Call.Fun.(*ast.SelectorExpr).Sel.NamePos, "identifier is GoMock Finish")
							}
							fmt.Println(reflect.TypeOf(i.(*ast.DeferStmt).Call.Args))
						}
						// log.Print(k, i)
					}
				}

		}
	})

	return nil, nil
}
