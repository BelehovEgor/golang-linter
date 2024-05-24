package handleTypeAssertionFailures

import "go/ast"

func do(i ast.Node) {
	_ = i.(*ast.ArrayType)

	_, ok := i.(*ast.BadDecl)
	if !ok {
		panic("awda")
	}
}
