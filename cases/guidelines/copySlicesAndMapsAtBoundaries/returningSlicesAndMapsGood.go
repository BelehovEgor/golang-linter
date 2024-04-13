package copySlicesAndMapsAtBoundaries

import "sync"

type StatsG struct {
	mu       sync.Mutex
	counters map[string]int
}

func (s *StatsG) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}

func makeSnapshotGood() {
	stats := Stats{counters: map[string]int{"a": 1, "b": 3}}
	snapshot := stats.Snapshot()
	print(snapshot) //just for usage variable
}

// Snapshot is now a copy.
