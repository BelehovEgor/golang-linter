package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"os"
)

func main2() {
	print(importUsage2())
}

func wrong_name_style(b bool) int {
	var a int
	var c = 2
	var d = 3
	area := Area(2.1)
	vol := Volume(3.1)
	if b {
		area += area
		vol += vol
		c++
		a = 100
	} else {
		d--
		a = 10
	}
	return a
}

func importUsage2() int {
	var b = jsoniter.ConfigCompatibleWithStandardLibrary
	fmt.Println("use")
	_, err := os.Hostname()
	if err != nil {
		b.Valid([]byte("{a:b}"))
		return 1
	}
	return 2
}

type Area2 float64
type Volume2 float64
