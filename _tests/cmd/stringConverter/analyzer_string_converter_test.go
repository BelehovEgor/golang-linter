package main

import (
	analyzer_string_converter "golang-linter/pkg/stringConverter"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_analyzer_string_converter(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "performance", "strConverter")
	analysistest.Run(t, testdata, analyzer_string_converter.Analyzer)
}
