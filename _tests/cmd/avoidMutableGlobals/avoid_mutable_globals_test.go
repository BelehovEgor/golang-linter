package main

import (
	avoid_mutable_globals "golang-linter/pkg/avoidMutableGlobals"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_avoid_mutable_globals(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "avoidMutableGlobals")
	analysistest.Run(t, testdata, avoid_mutable_globals.Analyzer)
}
