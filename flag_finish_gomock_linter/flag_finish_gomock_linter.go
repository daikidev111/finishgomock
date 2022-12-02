package flag_finish_gomock_linter

import (
	"go/ast"
	"log"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "flag_finish_gomock_linter is a linter that detects GoMock Finish Call when the testing package is used"

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
    // fset := token.NewFileSet()
    // f, err := parser.ParseFile(fset, "./testdata/src/a/a.go", nil, 0)
	// if err != nil {
	// 	panic(err)
	// }
    // ast.Print(fset, f)

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	go_mock_use_flag := false

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch callExprTyp := n.(type) {
		case *ast.CallExpr:
			// ast.Print(fset, n)
			go_mock_use_finish := false

			// new_cont_flag := false
			if (strings.Contains(pass.TypesInfo.TypeOf(callExprTyp).String(), "github.com/golang/mock/gomock.Controller")) {
				go_mock_use_flag = true
				// if (strings.Compare(callExprTyp.Fun.(*ast.SelectorExpr).Sel.Name, "NewController") == 0) {
				// 	new_cont_flag = true
				// }
			}

			if callExprTyp.Fun == nil {
				return
			}
			target_node, ok := callExprTyp.Fun.(*ast.SelectorExpr)
			if !ok {
				break
			}
			if (target_node.Sel.Name == "Finish") {
				go_mock_use_finish = true
			}

			if (go_mock_use_finish && go_mock_use_flag) {
				log.Print(target_node.Sel.NamePos, target_node.Sel.Name)
				pass.Reportf(target_node.Sel.NamePos, "identifier is GoMock Finish")
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
