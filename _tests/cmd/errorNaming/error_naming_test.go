package main

import (
	error_naming "golang-linter/pkg/errorNaming"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_error_naming(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "errorNaming")
	analysistest.Run(t, testdata, error_naming.Analyzer)
}
