package main

import (
	analyzer_string_converter "golang-linter/pkg/string-converter"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer_string_converter.Analyzer)
}
