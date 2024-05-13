package main

var _someData = 123

func sign(msg string) string {
	// Ought to be enough for anybody!
	_ = make(chan int, 64)

	// Size of one
	_ = make(chan int, 1) // or
	// Unbuffered channel, size of zero
	_ = make(chan int)

	x := 10
	_ = make(chan int, x)

	if _someData > 10 {
		return msg + "123"
	}

	return msg
}
