package main

import (
	channel_size "golang-linter/pkg/channel-size"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(channel_size.Analyzer)
}
