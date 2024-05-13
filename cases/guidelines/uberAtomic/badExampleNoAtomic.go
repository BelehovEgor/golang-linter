package uberAtomic

import "sync/atomic"

type fooBad struct {
	running int32 // atomic
}

func (f *fooBad) startBad() {
	if atomic.SwapInt32(&f.running, 1) == 1 {
		// already running…
		return
	}
	// start the Foo
}

func (f *fooBad) isRunningBad() bool {
	return f.running == 1 // race!
}
