package main

import (
	"fmt"
)

type Operation int

const (
	Add Operation = iota
	Subtract
	Multiply
)

type Operation1 int

const (
	Add1 Operation1 = iota + 1
	Subtract1
	Multiply1
)

func do() {
	fmt.Printf("Hello, Golang\n")
}
