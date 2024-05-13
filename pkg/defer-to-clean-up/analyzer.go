package defer_to_clean_up

import (
	"go/ast"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"

	"slices"
)

var Analyzer = &analysis.Analyzer{
	Name: "defertocleanup",
	Doc:  "Use defer to clean up resources such as files and locks.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	// TODO map type to func
	funcNames := []string{"Unlock", "Close"}

	inspect := func(node ast.Node) bool {
		_, ok := node.(*ast.DeferStmt)
		if ok {
			return false
		}

		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		// TODO save context for variables and check type of variable
		if slices.Contains(funcNames, selectorExpr.Sel.Name) {
			diagnostic := analysis.Diagnostic{
				Pos:      node.Pos(),
				End:      node.End(),
				Category: common.Warning,
				Message:  "Use defer to clean up resources such as files and locks.",
			}
			pass.Report(diagnostic)
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
