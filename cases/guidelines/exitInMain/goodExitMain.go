package exitInMain

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	body, err := readFileGood("path")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)
}

func readFileGood(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
