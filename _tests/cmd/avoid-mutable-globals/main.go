package main

import (
	avoid_mutable_globals "golang-linter/pkg/avoid-mutable-globals"
	"golang.org/x/tools/go/analysis/singlechecker"
	"os"
)

func main() {
	os.Args = []string{"avoidmutableglobals", "_tests/testdata/avoidMutableGlobals"}
	singlechecker.Main(avoid_mutable_globals.Analyzer)
}
