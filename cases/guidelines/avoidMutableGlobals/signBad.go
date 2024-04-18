package avoidMutableGlobals

import "time"

// sign.go

var _timeNow = time.Now

func sign(msg string) string {
	now := _timeNow()
	return signWithTime(msg, now)
}
