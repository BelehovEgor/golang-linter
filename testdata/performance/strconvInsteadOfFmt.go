package performance

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
		s := fmt.Sprint(rand.Int())
		print(s)
	}
}
