package deferToCleanUp

import (
	"fmt"
	"sync"
)

var counter int = 0 //  общий ресурс
func do() {

	ch := make(chan bool) // канал
	var mutex sync.Mutex  // определяем мьютекс
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The End")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() // блокируем доступ к переменной counter
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock() // want "Use defer to clean up resources such as files and locks."
	ch <- true
}

func work_but_better(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()         // блокируем доступ к переменной counter
	defer mutex.Unlock() // деблокируем доступ

	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}

	ch <- true
}
