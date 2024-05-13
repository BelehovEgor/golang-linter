package main

import (
	avoid_repeated_string_to_byte "golang-linter/pkg/avoid-repeated-string-to-byte"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(avoid_repeated_string_to_byte.Analyzer)
}
