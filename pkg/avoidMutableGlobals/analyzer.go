package avoidMutableGlobals

import (
	"fmt"
	"go/ast"

	"golang-linter/pkg/common"

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
			if element.Kind == ast.Var {

				diagnostic := analysis.Diagnostic{
					Pos:      node.Pos(),
					End:      node.End(),
					Category: common.Warning,
					Message:  fmt.Sprintf("Don't use mutable global variable '%s', use dependency injection.", element.Name),
				}
				pass.Report(diagnostic)
			}
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
