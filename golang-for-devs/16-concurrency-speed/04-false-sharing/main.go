package falsesharing

import (
	"sync"
)

/////////////////////////////////////////////////////////////
// False Sharing
/////////////////////////////////////////////////////////////

// CPU Caches:
// core -> cache -> [cache line - val1, val2] | [cache line - val3, val4]
// core -> cache -> [cache line - val1, val2]

//go:noinline
func inc(p *int64) {
	*p = *p + 1
}

type Counter struct {
	count int64
}

type PaddedCounter struct {
	count int64
	_     [128 - 8]byte
}

func CountSequential(numCounters int, countTo int) []Counter {
	counters := make([]Counter, numCounters)

	for i := 0; i < numCounters; i++ {
		c := &counters[i]
		for i := 0; i < countTo; i++ {
			inc(&c.count)
		}
	}

	return counters
}

func CountConcurrent(numCounters int, countTo int) []Counter {
	counters := make([]Counter, numCounters)

	wg := sync.WaitGroup{}

	for i := 0; i < numCounters; i++ {
		c := &counters[i]
		wg.Go(func() {
			for i := 0; i < countTo; i++ {
				inc(&c.count)
			}
		})
	}

	wg.Wait()

	return counters
}

func CountPadded(numCounters int, countTo int) []PaddedCounter {
	counters := make([]PaddedCounter, numCounters)

	wg := sync.WaitGroup{}

	for i := 0; i < numCounters; i++ {
		c := &counters[i]
		wg.Go(func() {
			for i := 0; i < countTo; i++ {
				inc(&c.count)
			}
		})
	}

	wg.Wait()

	return counters
}
