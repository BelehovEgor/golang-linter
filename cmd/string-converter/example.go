package main

import (
	"fmt"
	"log"
)

func myLog(format string, args ...interface{}) {
	log.Printf(format, args...)
	fmt.Sprint(123)
}
