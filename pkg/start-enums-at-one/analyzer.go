package start_enums_at_one

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "startenumsatone",
	Doc: "The standard way of introducing enumerations in Go is to declare a custom type and a const group with iota." +
		"Since variables have a 0 default value, you should usually start your enums on a non-zero value.",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		genDecl, ok := node.(*ast.GenDecl)
		if !ok {
			return true
		}

		if genDecl.Tok != token.CONST {
			return true
		}

		if len(genDecl.Specs) == 0 {
			return true
		}

		firstSpecElement, ok := genDecl.Specs[0].(*ast.ValueSpec)
		if !ok {
			return true
		}

		if len(firstSpecElement.Names) != 1 || len(firstSpecElement.Values) != 1 {
			return true
		}
		nameIdent := firstSpecElement.Names[0]
		value := firstSpecElement.Values[0]

		typeIdent, ok := firstSpecElement.Type.(*ast.Ident)
		if !ok {
			return true
		}

		valueIdent, ok := value.(*ast.Ident)
		if !ok {
			return true
		}

		if valueIdent.Name != "iota" {
			return true
		}

		diagnostic := analysis.Diagnostic{
			Pos:      node.Pos(),
			End:      node.End(),
			Category: common.Notion,
			Message:  fmt.Sprintf("In a constant '%s', the starting value '%s' is equal 0.", typeIdent.Name, nameIdent.Name),
		}
		pass.Report(diagnostic)

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
