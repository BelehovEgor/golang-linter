package main

import (
	start_enums_at_one "golang-linter/pkg/start-enums-at-one"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(start_enums_at_one.Analyzer)
}
