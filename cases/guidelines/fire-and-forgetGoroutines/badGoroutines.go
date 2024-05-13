package fire_and_forgetGoroutines

import "time"

func startGogB() {
	go func() {
		for {
			print("aaaa")
			time.Sleep(20)
		}
	}()
}
