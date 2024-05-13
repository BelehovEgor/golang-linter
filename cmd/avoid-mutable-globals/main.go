package main

import (
	avoid_mutable_globals "golang-linter/pkg/avoid-mutable-globals"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(avoid_mutable_globals.Analyzer)
}
