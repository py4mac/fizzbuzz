package repository

import (
	"sort"

	"github.com/py4mac/fizzbuzz/pkg/fizzbuzz"
	"github.com/py4mac/fizzbuzz/pkg/stats"
)

type statsInMemory struct {
	stats map[fizzbuzz.Fizzbuz]int32
}

func NewStatsInMemory() stats.Stats {
	s := &statsInMemory{}
	s.stats = make(map[fizzbuzz.Fizzbuz]int32)

	return s
}

func (s *statsInMemory) Record(e fizzbuzz.Fizzbuz) error {
	_, ok := s.stats[e]
	if !ok {
		s.stats[e] = 1
		return nil
	}
	s.stats[e]++

	return nil
}

func (s *statsInMemory) Process() (stats.Statistics, error) {
	keys := make([]fizzbuzz.Fizzbuz, 0, len(s.stats))
	for key := range s.stats {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return s.stats[keys[i]] > s.stats[keys[j]] })

	return stats.Statistics{
		Hits:    s.stats[keys[0]],
		Fizzbuz: keys[0],
	}, nil
}
