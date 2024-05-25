package performance

import (
	"os"
)

func badExampleMapCapacity() {
	m := make(map[string]os.DirEntry)

	files, _ := os.ReadDir("./files")
	for _, f := range files {
		m[f.Name()] = f
	}
}

func goodExampleMapCapacity() {
	files, _ := os.ReadDir("./files")

	m := make(map[string]os.DirEntry, len(files))
	for _, f := range files {
		m[f.Name()] = f
	}
}
