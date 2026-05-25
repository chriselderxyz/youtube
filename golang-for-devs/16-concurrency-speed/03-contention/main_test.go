package contention

import (
	"concurrency-speed/helpers"
	"fmt"
	"testing"
)

var sink int64

func BenchmarkContention(b *testing.B) {
	
	nums := helpers.RandomSlice[int64](10_000)
	chunkSizes := []int{5_000, 1_000, 500, 100, 50, 5, 2}

	for _, size := range chunkSizes {

		b.Run(fmt.Sprintf("SumConcurrent-num-routines-%d", len(nums)/size), func(b *testing.B) {
			for b.Loop() {
				sink = SumConcurrent(nums, size)
			}
		})

		b.Run(fmt.Sprintf("SumAtomic-num-routines-%d", len(nums)/size), func(b *testing.B) {
			for b.Loop() {
				sink = SumAtomic(nums, size)
			}
		})

		b.Run(fmt.Sprintf("SumLock-num-routines-%d", len(nums)/size), func(b *testing.B) {
			for b.Loop() {
				sink = SumLock(nums, size)
			}
		})

		fmt.Println("====================================================")
	}
}
