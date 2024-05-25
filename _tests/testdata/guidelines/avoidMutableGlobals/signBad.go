package avoidMutableGlobals // want  "Don't use mutable global variable '_timeNow', use dependency injection."

import "time"

// sign.go

var _timeNow = time.Now

func sign(msg string) string {
	now := _timeNow()
	return signWithTime(msg, now)
}
