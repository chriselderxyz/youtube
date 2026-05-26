package diminishingreturns

import (
	"sync"
)

func Work(x int64) int64 {
	for i := int64(1); i <= 200; i++ {
		x = x*1_103_515_245 + i*97 + 12_345
		x = x*x + 3*x + i + 7
		x = x/3 + x*5 + 11
	}

	return x
}

func SumSequential(nums []int64) int64 {
	total := int64(0)
	for _, n := range nums {
		total += Work(n)
	}

	return total
}

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
				total += Work(nums[i])
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
