package contention

import (
	"sync"
	"sync/atomic"
)

func SumConcurrent(nums []int64, chunkSize int) int64 {
	wg := sync.WaitGroup{}
	numChunks := (len(nums) + chunkSize - 1) / chunkSize
	chunkTotal := make([]int64, numChunks)

	chunk := 0

	for start := 0; start < len(nums); start += chunkSize {
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}

		chunk := chunk

		wg.Go(func() {
			total := int64(0)
			for i := start; i < end; i++ {
				total += nums[i]
			}

			chunkTotal[chunk] = total
		})

		chunk++
	}

	wg.Wait()

	total := int64(0)
	for _, n := range chunkTotal {
		total = total + n
	}

	return total
}

func SumAtomic(nums []int64, chunkSize int) int64 {
	wg := sync.WaitGroup{}
	total := int64(0)

	for start := 0; start < len(nums); start += chunkSize {
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}

		wg.Go(func() {
			for i := start; i < end; i++ {
				atomic.AddInt64(&total, nums[i])
			}
		})
	}

	wg.Wait()

	return total
}

func SumLock(nums []int64, chunkSize int) int64 {
	wg := sync.WaitGroup{}
	total := int64(0)
	mu := sync.Mutex{}

	for start := 0; start < len(nums); start += chunkSize {
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}

		wg.Go(func() {
			for i := start; i < end; i++ {
				mu.Lock()
				total += nums[i]
				mu.Unlock()
			}
		})
	}

	wg.Wait()
	return total
}
