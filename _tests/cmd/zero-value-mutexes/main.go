package main

import (
	zero_value_mutexes "golang-linter/pkg/zero-value-mutexes"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(zero_value_mutexes.Analyzer)
}
