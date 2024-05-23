package copySlicesAndMapsAtBoundaries

import "sync"

type Stats struct {
	mu       sync.Mutex
	counters map[string]int
}

// Snapshot returns the current stats.
func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counters
}

// snapshot is no longer protected by the mutex, so any
// access to the snapshot is subject to data races.
func makeSnapshotBad() {
	stats := Stats{counters: map[string]int{"a": 1, "b": 3}}
	snapshot := stats.Snapshot()
	print(snapshot) //just for usage variable
}
