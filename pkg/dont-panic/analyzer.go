package dont_panic

import (
	"go/ast"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "dontpanic",
	Doc:  "Code running in production must avoid panics.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		funIdent, ok := callExpr.Fun.(*ast.Ident)
		if !ok || funIdent.Name != "panic" {
			return true
		}

		diagnostic := analysis.Diagnostic{
			Pos:      node.Pos(),
			End:      node.End(),
			Category: common.Notion,
			Message:  "If an error occurs, the function must return an error and allow the caller to decide how to handle it.",
		}
		pass.Report(diagnostic)

		return false
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
