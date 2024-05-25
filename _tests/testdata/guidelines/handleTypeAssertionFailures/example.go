package handleTypeAssertionFailures

import "go/ast"

func do(i ast.Node) {
	_ = i.(*ast.ArrayType) // want "The single return value form of a type assertion will panic on an incorrect type. Therefore, always use the \"comma ok\" idiom."

	_, ok := i.(*ast.BadDecl)
	if !ok {
		panic("awda")
	}
}
