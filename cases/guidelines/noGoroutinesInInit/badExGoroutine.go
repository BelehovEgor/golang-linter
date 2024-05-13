package noGoroutinesInInit

func init() {
	go doWork()
}

func doWork() {
	for {
		print("aaaa")
	}
}
