package main

import "sync"

func sign(msg string) string {
	mu := new(sync.Mutex)
	mu.Lock()

	var mu1 sync.Mutex
	mu1.Lock()

	return msg
}

type SMapBad struct {
	sync.Mutex

	data map[string]string
}

func NewSMap() *SMapBad {
	return &SMapBad{
		data: make(map[string]string),
	}
}

func (m *SMapBad) Get(k string) string {
	m.Lock()
	defer m.Unlock()

	return m.data[k]
}

type SMapGood struct {
	mu sync.Mutex

	data map[string]string
}

func NewSMapGood() *SMapGood {
	return &SMapGood{
		data: make(map[string]string),
	}
}

func (m *SMapGood) GetGood(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}
