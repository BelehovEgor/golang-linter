package main

import (
	zero_value_mutexes "github.com/BelehovEgor/golang-linter/pkg/zeroValueMutexes"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_zero_value_mutexes(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "zeroValueMutex")
	analysistest.Run(t, testdata, zero_value_mutexes.Analyzer)
}
