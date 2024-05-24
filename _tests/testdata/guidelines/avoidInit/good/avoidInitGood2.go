package good

import (
	"os"
	"path"
)

type Config struct {
	// ...
}

func loadConfig() Config {
	cwd, err := os.Getwd()
	// handle err
	print(err)
	raw, err := os.ReadFile(
		path.Join(cwd, "config", "config.yaml"),
	)
	print(raw)
	// handle err

	var config Config
	config = Config{}

	return config
}
