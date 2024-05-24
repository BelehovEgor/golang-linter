package uberAtomic

import "sync/atomic"

type fooGood struct {
	running atomic.Bool
}

func (f *fooGood) startGood() {
	if f.running.Swap(true) {
		// already running…
		return
	}
	// start the Foo
}

func (f *fooGood) isRunningGood() bool {
	return f.running.Load()
}
