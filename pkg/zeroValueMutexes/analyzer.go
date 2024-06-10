package zeroValueMutexes

import (
	"go/ast"

	"github.com/BelehovEgor/golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

// Analyzer for golang-linter
var Analyzer = &analysis.Analyzer{
	Name: "zerovaluemutexes",
	Doc:  "The zero-value of sync.Mutex and sync.RWMutex is valid, so you almost never need a pointer to a mutex.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectCreatingMutex := func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		fun, ok := callExpr.Fun.(*ast.Ident)
		if !ok || fun.Name != "new" {
			return true
		}

		if len(callExpr.Args) != 1 {
			return true
		}

		if !isMutex(callExpr.Args[0]) {
			return true
		}

		diagnostic := analysis.Diagnostic{
			Pos:      node.Pos(),
			End:      node.End(),
			Category: common.Warning,
			Message:  "The zero-value of sync.Mutex and sync.RWMutex is valid, so you almost never need a pointer to a mutex.",
		}
		pass.Report(diagnostic)

		return true
	}

	inspectStructWithMutex := func(node ast.Node) bool {
		structFieldList, ok := node.(*ast.FieldList)
		if !ok {
			return true
		}

		for _, field := range structFieldList.List {
			if len(field.Names) != 0 {
				continue
			}

			if !isMutex(field.Type) {
				return true
			}

			diagnostic := analysis.Diagnostic{
				Pos:      field.Pos(),
				End:      field.End(),
				Category: common.Warning,
				Message:  "Do not embed the mutex on the struct, even if the struct is not exported.",
			}
			pass.Report(diagnostic)
		}

		return true
	}

	inspect := func(node ast.Node) bool {
		return inspectCreatingMutex(node) && inspectStructWithMutex(node)
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}

func isMutex(expr ast.Expr) bool {
	mutexSelectorExpr, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	xIdent, ok := mutexSelectorExpr.X.(*ast.Ident)
	if !ok || xIdent.Name != "sync" {
		return false
	}

	return mutexSelectorExpr.Sel.Name == "Mutex" || mutexSelectorExpr.Sel.Name == "RWMutex"
}
