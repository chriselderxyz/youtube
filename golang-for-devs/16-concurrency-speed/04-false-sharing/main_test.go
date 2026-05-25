package falsesharing

import (
	"fmt"
	"runtime"
	"testing"
)

var sink1 []Counter
var sink2 []PaddedCounter

func BenchmarkFalseSharing(b *testing.B) {

	countTo := 1_000_000
	numCounters := runtime.GOMAXPROCS(0)

	b.Run("CountSequential", func(b *testing.B) {
		for b.Loop() {
			sink1 = CountSequential(numCounters, countTo)
		}
	})

	b.Run("CountConcurrent", func(b *testing.B) {
		for b.Loop() {
			sink1 = CountConcurrent(numCounters, countTo)
		}
	})

	b.Run("CountPadded-%d", func(b *testing.B) {
		for b.Loop() {
			sink2 = CountPadded(numCounters, countTo)
		}
	})

	fmt.Println("================================================")

}
