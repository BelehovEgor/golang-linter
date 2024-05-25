package main

import (
	defer_to_clean_up "golang-linter/pkg/deferToCleanUp"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_defer_to_clean_up(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "deferToCleanUp")
	analysistest.Run(t, testdata, defer_to_clean_up.Analyzer)
}
