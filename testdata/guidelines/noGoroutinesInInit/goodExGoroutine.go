package noGoroutinesInInit

type Worker struct {
	stop, done chan struct{}
}

func NewWorker() *Worker {
	w := &Worker{
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
	go w.doWork()
	return w
}

func (w *Worker) doWork() {
	defer close(w.done)
	for {
		select {
		case <-w.stop:
			return
		}
	}
}

// Shutdown tells the worker to stop
// and waits until it has finished.
func (w *Worker) Shutdown() {
	close(w.stop)
	<-w.done
}
