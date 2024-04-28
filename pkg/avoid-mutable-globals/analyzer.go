package avoid_mutable_globals

import (
	"go/ast"
	ObjKind "go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "avoidmutableglobals",
	Doc:  "Avoid mutating global variables, instead opting for dependency injection.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		file, ok := node.(*ast.File)
		if !ok {
			return true
		}

		objects := file.Scope.Objects
		for _, element := range objects {
			if element.Kind == ObjKind.Var {
				pass.Reportf(node.Pos(), "Don't use mutable global variable '%s', use dependency injection.", element.Name)
			}
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
