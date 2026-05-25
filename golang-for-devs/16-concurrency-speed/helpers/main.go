package helpers

import "math/rand/v2"

func main() {

}

type Ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func RandomSlice[T Ints](size int) []T {
	in := make([]T, size)
	for i := range size {
		in[i] = T(rand.IntN(1000) + 1)
	}

	return in
}
