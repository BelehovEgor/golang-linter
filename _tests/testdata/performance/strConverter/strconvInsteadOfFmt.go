package strConverter

import (
	"fmt"
	"math/rand"
	"strconv"
)

func goodExampleStrconv() {
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(rand.Int())
		print(s)
	}
}

func badExampleFmt() {
	for i := 0; i < 10; i++ {
		s := fmt.Sprint(rand.Int()) // want "Use 'strconv' instead 'fmt' for converting type to/from string"
		print(s)
	}
}
