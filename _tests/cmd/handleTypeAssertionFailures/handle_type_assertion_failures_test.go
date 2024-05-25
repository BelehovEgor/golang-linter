package main

import (
	handle_type_assertion_failures "golang-linter/pkg/handleTypeAssertionFailures"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_handle_type_assertion_failures(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "handleTypeAssertionFailures")
	analysistest.Run(t, testdata, handle_type_assertion_failures.Analyzer)
}
