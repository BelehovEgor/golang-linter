package main

import (
	start_enums_at_one "golang-linter/pkg/startEnumsAtOne"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_start_enums_at_one(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "enumsWithOne")
	analysistest.Run(t, testdata, start_enums_at_one.Analyzer)
}
