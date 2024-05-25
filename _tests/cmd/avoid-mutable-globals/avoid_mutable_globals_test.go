package main

import (
	avoid_mutable_globals "golang-linter/pkg/avoid-mutable-globals"
	"golang.org/x/tools/go/analysis/analysistest"
	"log"
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
	log.Printf(testdata)
	analysistest.Run(t, testdata, avoid_mutable_globals.Analyzer)
}
