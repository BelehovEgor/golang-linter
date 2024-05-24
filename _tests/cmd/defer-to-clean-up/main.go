package main

import (
	defer_to_clean_up "golang-linter/pkg/defer-to-clean-up"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(defer_to_clean_up.Analyzer)
}
