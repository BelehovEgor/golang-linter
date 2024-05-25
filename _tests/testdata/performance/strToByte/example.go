package strToByte

import (
	"fmt"
)

func do_smth() {
	fmt.Printf("Hello, Golang\n")

	_ = []byte("Hello world")

	for i := 0; i < 10; i++ {
		fmt.Print([]byte("Hello world")) // want "Do not create byte slices from a fixed string repeatedly. Instead, perform the conversion once and capture the result."
		_ = []byte("Hello world")        // want "Do not create byte slices from a fixed string repeatedly. Instead, perform the conversion once and capture the result."
	}
}
