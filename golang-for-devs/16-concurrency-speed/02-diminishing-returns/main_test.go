package diminishingreturns

import (
	"concurrency-speed/helpers"
	"fmt"
	"testing"
)

var sink int64

func BenchmarkDiminishingReturns(b *testing.B) {

	nums := helpers.RandomSlice[int64](10_000)

	b.Run("SumSequential: ", func(b *testing.B) {
		for b.Loop() {
			sink = SumSequential(nums)
		}
	})

	fmt.Println("====================================================")

	chunkSizes := []int{5_000, 1_000, 500, 100, 50, 25, 10}

	for _, size := range chunkSizes {
		b.Run(fmt.Sprintf("SumConcurrent-num-routines-%d", len(nums)/size), func(b *testing.B) {
			for b.Loop() {
				sink = SumConcurrent(nums, size)
			}
		})
	}
}
