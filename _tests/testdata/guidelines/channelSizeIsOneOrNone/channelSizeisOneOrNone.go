package channelSizeIsOneOrNone

func channelSizeIsOneOrNoneBadEx() {
	// Ought to be enough for anybody!
	c := make(chan int, 64) // want "Channels should usually have a size of one or be unbuffered. Any other size must be subject to a high level of scrutiny"
	print(len(c))
}

func channelSizeIsOneOrNoneGoodEx() {
	// Ought to be enough for anybody!
	// Size of one
	c := make(chan int, 1) // or
	// Unbuffered channel, size of zero
	d := make(chan int)
	print(len(c))
	print(len(d))
}
