package avoidRepeatedStringToByte

import (
	"go/ast"
	"go/token"

	"golang-linter/pkg/common"

	"golang.org/x/tools/go/analysis"
)

// Analyzer for golang-linter
var Analyzer = &analysis.Analyzer{
	Name: "avoidrepeatedstringtobyte",
	Doc:  "Avoid creating byte slices from a fixed string repeatedly.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		forStmt, ok := node.(*ast.ForStmt)
		if !ok {
			return true
		}

		return processExprStmt(pass, forStmt.Body)
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}

func processExprStmt(pass *analysis.Pass, node ast.Node) bool {
	blockStmt, ok := node.(*ast.BlockStmt)
	if ok {
		for _, arg := range blockStmt.List {
			processExprStmt(pass, arg)
		}

		return false
	}

	exrStmt, ok := node.(*ast.ExprStmt)
	if ok {
		return processExprStmt(pass, exrStmt.X)
	}

	assignStmt, ok := node.(*ast.AssignStmt)
	if ok {
		for _, arg := range assignStmt.Rhs {
			processExprStmt(pass, arg)
		}

		return false
	}

	callExpr, ok := node.(*ast.CallExpr)
	if !ok {
		return true
	}

	arrayType, ok := callExpr.Fun.(*ast.ArrayType)
	if !ok {
		for _, arg := range callExpr.Args {
			processExprStmt(pass, arg)
		}

		return false
	}

	identArrayType, ok := arrayType.Elt.(*ast.Ident)
	if !ok || identArrayType.Name != "byte" {
		return true
	}

	if len(callExpr.Args) != 1 {
		return true
	}

	basicLit, ok := callExpr.Args[0].(*ast.BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		return true
	}

	diagnostic := analysis.Diagnostic{
		Pos:      node.Pos(),
		End:      node.End(),
		Category: common.Warning,
		Message:  "Do not create byte slices from a fixed string repeatedly. Instead, perform the conversion once and capture the result.",
	}
	pass.Report(diagnostic)

	return false
}
