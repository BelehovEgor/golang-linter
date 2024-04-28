package zeroValueMutex

import "sync"

type smap struct {
	sync.Mutex // only for unexported types

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
	mu := new(sync.Mutex)
	mu.Lock()
}
