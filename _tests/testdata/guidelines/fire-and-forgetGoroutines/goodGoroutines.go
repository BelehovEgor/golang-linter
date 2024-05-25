package fire_and_forgetGoroutines

import "time"

var (
	stop = make(chan struct{}) // tells the goroutine to stop
	done = make(chan struct{}) // tells us that the goroutine exited
)

func runGorG() {
	go func() {
		defer close(done)

		ticker := time.NewTicker(20)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				print("aaaa")
			case <-stop:
				return
			}
		}
	}()
}

func stopGorG() {
	close(stop) // signal the goroutine to stop
	<-done      // and wait for it to exit
}
