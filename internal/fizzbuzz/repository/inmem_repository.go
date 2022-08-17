package repository

import (
	"context"
	"sort"
	"sync"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

type fbInMemory struct {
	mem  map[domain.Fizzbuz]int32
	lock sync.RWMutex
}

func NewFBInMemory() fizzbuzz.Repository {
	s := &fbInMemory{}
	s.mem = make(map[domain.Fizzbuz]int32)

	return s
}

func (f *fbInMemory) Record(ctx context.Context, e domain.Fizzbuz) ([]string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, ok := f.mem[e]

	if !ok {
		f.mem[e] = 1
		return e.Process(ctx)
	}
	f.mem[e]++

	return e.Process(ctx)
}

func (f *fbInMemory) Process(ctx context.Context) (*domain.Statistics, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	keys := make([]domain.Fizzbuz, 0, len(f.mem))

	for key := range f.mem {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return f.mem[keys[i]] > f.mem[keys[j]] })

	return &domain.Statistics{
		Hits:    f.mem[keys[0]],
		Fizzbuz: keys[0],
	}, nil
}
