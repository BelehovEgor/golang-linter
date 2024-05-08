package main

import (
	"fmt"
)

func do_smth() {
	fmt.Printf("Hello, Golang\n")

	_ = []byte("Hello world")

	for i := 0; i < 10; i++ {
		fmt.Print([]byte("Hello world"))
		_ = []byte("Hello world")
	}
}
