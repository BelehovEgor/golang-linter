package main

import (
	dont_panic "golang-linter/pkg/dont-panic"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(dont_panic.Analyzer)
}
