package exitInMain

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	body := readFileBad("path")
	fmt.Println(body)
}

func readFileBad(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
