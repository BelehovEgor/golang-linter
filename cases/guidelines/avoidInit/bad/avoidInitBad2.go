package bad

import (
	"os"
	"path"
)

type Config struct {
	// ...
}

var _config Config

func init() {
	// Bad: based on current directory
	cwd, _ := os.Getwd()

	// Bad: I/O
	raw, _ := os.ReadFile(
		path.Join(cwd, "config", "config.yaml"),
	)
	print(raw)

	_config = Config{}
}
