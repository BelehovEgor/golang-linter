package zeroValueMutex

import "sync"

type smap struct {
	sync.Mutex // want "Do not embed the mutex on the struct, even if the struct is not exported."
	// only for unexported types

	data map[string]string
}

func newSMap() *smap {
	return &smap{
		data: make(map[string]string),
	}
}

func (m *smap) GetBadMutex(k string) string {
	m.Lock()
	defer m.Unlock()

	return m.data[k]
}

func badExampleMutex() {
	mu := new(sync.Mutex) // want "The zero-value of sync.Mutex and sync.RWMutex is valid, so you almost never need a pointer to a mutex."
	mu.Lock()
}
