package bad

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("missing file")
	}
	name := args[0]

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// If we call log.Fatal after this line,
	// f.Close will not be called.

	b, err := io.ReadAll(f)
	if err != nil {
		os.Exit(1)
	}
	print(b)

	// ...
}
