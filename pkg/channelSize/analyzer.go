package channelSize

import (
	"go/ast"
	"go/token"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

// Analyzer for golang-linter
var Analyzer = &analysis.Analyzer{
	Name: "channelsize",
	Doc:  "Channels should usually have a size of one or be unbuffered.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		funIdent, ok := callExpr.Fun.(*ast.Ident)
		if !ok {
			return true
		}

		if funIdent.Name != "make" {
			return true
		}

		if len(callExpr.Args) != 2 {
			return true
		}

		_, ok = callExpr.Args[0].(*ast.ChanType)
		if !ok {
			return true
		}

		constChanSize, ok := callExpr.Args[1].(*ast.BasicLit)
		if !ok || constChanSize.Kind != token.INT {
			return true
		}

		if constChanSize.Value == "1" {
			return true
		}

		diagnostic := analysis.Diagnostic{
			Pos:      node.Pos(),
			End:      node.End(),
			Category: common.Notion,
			Message:  "Channels should usually have a size of one or be unbuffered. Any other size must be subject to a high level of scrutiny",
		}
		pass.Report(diagnostic)

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
