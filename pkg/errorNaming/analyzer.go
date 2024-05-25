package errorNaming

import (
	"go/ast"
	"go/token"
	"strings"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "errornaming",
	Doc:  "For error values stored as global variables, use the prefix Err or err depending on whether they're exported.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		file, ok := node.(*ast.File)
		if !ok {
			return true
		}

		for _, decl := range file.Decls {
			varGenDecl, ok := decl.(*ast.GenDecl)
			if !ok || varGenDecl.Tok != token.VAR {
				continue
			}

			for _, spec := range varGenDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok || len(valueSpec.Names) != 1 || len(valueSpec.Values) != 1 {
					continue
				}

				value, ok := valueSpec.Values[0].(*ast.CallExpr)
				if !ok {
					continue
				}

				selectorExpr, ok := value.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}

				xIdent, ok := selectorExpr.X.(*ast.Ident)
				if !ok || xIdent.Name != "errors" || selectorExpr.Sel.Name != "New" {
					continue
				}

				variableName := valueSpec.Names[0].Name
				if strings.HasPrefix(variableName, "err") || strings.HasPrefix(variableName, "Err") {
					continue
				}

				diagnostic := analysis.Diagnostic{
					Pos:      valueSpec.Names[0].Pos(),
					End:      valueSpec.Names[0].End(),
					Category: common.Notion,
					Message:  "Global error variable has to start with 'err' or 'Err' prefix",
				}
				pass.Report(diagnostic)
			}
		}

		return false
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
