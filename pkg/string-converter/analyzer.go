package analyzer_string_converter

import (
	"go/ast"

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

		pass.Reportf(node.Pos(), "Use 'strconv' instead 'fmt' for converting type to/from string")

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
