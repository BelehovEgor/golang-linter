package avoidMutableGlobals

import (
	"fmt"
	"time"
)

// sign_test.go

func SignerTest() {
	s := newSigner()
	t2 := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	s.now = func() time.Time {
		return t2
	}

	fmt.Println("test"+t2.GoString() == s.Sign("test"))
}
