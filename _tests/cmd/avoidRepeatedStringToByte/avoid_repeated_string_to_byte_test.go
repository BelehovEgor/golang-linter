package main

import (
	avoid_repeated_string_to_byte "golang-linter/pkg/avoidRepeatedStringToByte"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_avoid_repeated_string_to_byte(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "performance", "strToByte")
	analysistest.Run(t, testdata, avoid_repeated_string_to_byte.Analyzer)
}
