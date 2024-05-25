package avoidMutableGlobals

import (
	"fmt"
	"time"
)

func SignTest() {
	oldTimeNow := _timeNow
	t2 := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	_timeNow = func() time.Time {
		return t2
	}
	defer func() { _timeNow = oldTimeNow }()

	fmt.Println("test"+t2.GoString() == sign("test"))
}
