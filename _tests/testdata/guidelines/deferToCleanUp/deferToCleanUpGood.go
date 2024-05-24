package deferToCleanUp

import "sync"

type deferEx struct {
	mu    sync.Mutex
	count int
}

func (d *deferEx) Lock() {
	d.mu.Lock()
}

func (d *deferEx) Unlock() {
	d.mu.Lock()
}

func deferToCleanUpGoodEx() int {
	p := deferEx{count: 0}
	p.Lock()
	defer p.Unlock()

	if p.count < 10 {
		return p.count
	}

	p.count++
	return p.count
	// more readable
}
