package zeroValueMutex

import "sync"

type SMap struct {
	mu sync.Mutex

	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) GetGoodMutex(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}

func goodExampleMutex() {
	var mu sync.Mutex
	mu.Lock()
}
