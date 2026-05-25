package smallpipeline

import (
	"sync"
)

func Run(in <-chan int, transform func(int) int, numWorkers int) <-chan int {
	wg := sync.WaitGroup{}
	out := make(chan int)

	for i := 0; i < numWorkers; i++ {
		wg.Go(func() {
			for n := range in {
				out <- transform(n)
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Square(n int) int {
	return n * n
}

func Double(n int) int {
	return n * 2
}

func AddSeven(n int) int {
	return n + 7
}

func PipelineSequential(in <-chan int) []int {
	out := []int{}

	for n := range in {
		val := Square(n)
		val = Double(val)
		val = AddSeven(val)

		out = append(out, val)
	}

	return out
}

func PipelineConcurrent(in <-chan int, numWorkers int) []int {
	out1 := Run(in, Square, numWorkers)
	out2 := Run(out1, Double, numWorkers)
	out3 := Run(out2, AddSeven, numWorkers)

	out := []int{}
	for n := range out3 {
		out = append(out, n)
	}

	return out
}
