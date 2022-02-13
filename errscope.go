package errscope

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "errscope is analyzing whether the scope of the error type can be narrowed"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "errscope",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.IfStmt)(nil),
	}

	assignStmtMap := make(map[*ast.AssignStmt]int)
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.IfStmt:
			a, ok := n.Init.(*ast.AssignStmt)
			if ok {
				assignStmtMap[a]++
			}
		case *ast.AssignStmt:
			if len(n.Lhs) != len(n.Rhs) {
				return
			}

			if _, ok := assignStmtMap[n]; ok {
				return
			}

			for _, expr := range n.Lhs {
				typ := pass.TypesInfo.TypeOf(expr)
				errType := types.Universe.Lookup("error").Type()
				if types.Identical(typ, errType) {
					pass.Reportf(n.Pos(), "this error type can be narrowed in scope")
				}
			}
		}
	})

	return nil, nil
}
