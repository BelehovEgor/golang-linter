package channelSizeIsOneOrNone

var _someData = 123

func sign(msg string) string {
	// Ought to be enough for anybody!
	_ = make(chan int, 64) // want "Channels should usually have a size of one or be unbuffered. Any other size must be subject to a high level of scrutiny"

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
