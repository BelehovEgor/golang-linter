package main

import (
	handle_type_assertion_failures "golang-linter/pkg/handle-type-assertion-failures"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(handle_type_assertion_failures.Analyzer)
}
