package analyzer_string_converter

import (
	"go/ast"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "stringconverter",
	Doc:  "Checks that using strconv for converting types to/from strings.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	convertToFromStringFmtFuncs := map[string]bool{
		"Sprint":   true,
		"Sprintf":  true,
		"Sprintln": true,
	}

	inspect := func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		fun, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		x, ok := fun.X.(*ast.Ident)
		if !ok || x.Name != "fmt" {
			return true
		}

		_, ok = convertToFromStringFmtFuncs[fun.Sel.Name]
		if !ok {
			return true
		}

		diagnostic := analysis.Diagnostic{
			Pos:      node.Pos(),
			End:      node.End(),
			Category: common.Error,
			Message:  "Use 'strconv' instead 'fmt' for converting type to/from string",
		}
		pass.Report(diagnostic)

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
