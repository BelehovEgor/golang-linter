package main

import (
	dont_panic "golang-linter/pkg/dont-panic"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_dont_panic(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "panics")
	analysistest.Run(t, testdata, dont_panic.Analyzer)
}
