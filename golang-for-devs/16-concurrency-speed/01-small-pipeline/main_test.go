package smallpipeline

import (
	"concurrency-speed/helpers"
	"fmt"
	"testing"
)

var sink []int

func GetChannel(inputs []int) chan int {
	ch := make(chan int)
	go func() {
		for _, n := range inputs {
			ch <- n
		}
		close(ch)
	}()

	return ch
}

func BenchmarkSmallPipeline(b *testing.B) {

	inputs := helpers.RandomSlice[int](1000)

	b.Run("PipelineSequential", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			ch := GetChannel(inputs)
			b.StartTimer()
			sink = PipelineSequential(ch)
		}
	})

	fmt.Println("===================================================")

	numWorkers := []int{1, 5, 10, 50, 100}

	for _, w := range numWorkers {
		b.Run(fmt.Sprintf("PipelineConcurrent-workers-%d", w), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				ch := GetChannel(inputs)
				b.StartTimer()
				sink = PipelineConcurrent(ch, w)
			}
		})
	}

}
