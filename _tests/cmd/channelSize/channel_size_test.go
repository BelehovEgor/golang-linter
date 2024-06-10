package main

import (
	channel_size "github.com/BelehovEgor/golang-linter/pkg/channelSize"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func Test_channel_size(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata", "guidelines", "channelSizeIsOneOrNone")
	analysistest.Run(t, testdata, channel_size.Analyzer)
}
