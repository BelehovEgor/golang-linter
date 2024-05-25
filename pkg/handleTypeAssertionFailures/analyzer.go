package handleTypeAssertionFailures

import (
	"go/ast"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

// Analyzer for golang-linter
var Analyzer = &analysis.Analyzer{
	Name: "handletypeassertionfailures",
	Doc:  "Always use the \"comma ok\" idiom.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		assignStmt, ok := node.(*ast.AssignStmt)
		if !ok || len(assignStmt.Rhs) != 1 {
			return true
		}

		_, ok = assignStmt.Rhs[0].(*ast.TypeAssertExpr)
		if !ok {
			return true
		}

		if len(assignStmt.Lhs) != 1 {
			return true
		}

		diagnostic := analysis.Diagnostic{
			Pos:      node.Pos(),
			End:      node.End(),
			Category: common.Notion,
			Message:  "The single return value form of a type assertion will panic on an incorrect type. Therefore, always use the \"comma ok\" idiom.",
		}
		pass.Report(diagnostic)

		return false
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
