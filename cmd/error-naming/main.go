package main

import (
	error_naming "golang-linter/pkg/error-naming"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(error_naming.Analyzer)
}
