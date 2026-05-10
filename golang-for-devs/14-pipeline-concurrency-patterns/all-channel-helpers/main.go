package channelhelpers

import (
	"context"
	"sync"
)

func CancelOrSend[T any](ctx context.Context, ch chan<- T, val T) bool {
	select {
	case <-ctx.Done():
		return false
	case ch <- val:
		return true
	}
}

func CancelOrReceive[T any](ctx context.Context, in <-chan T) (val T, ok bool) {
	select {
	case <-ctx.Done():
		return
	case val, ok = <-in:
		return
	}
}

func Merge[T any](ctx context.Context, buffer int, inputs ...<-chan T) <-chan T {
	out := make(chan T, buffer)

	wg := sync.WaitGroup{}

	for _, ch := range inputs {
		wg.Go(func() {
			if ch == nil {
				return
			}

			for {
				val, ok := CancelOrReceive(ctx, ch)
				if !ok {
					return
				}

				if !CancelOrSend(ctx, out, val) {
					return
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Split[T any](ctx context.Context, in <-chan T, numChannels int, buffer int) []<-chan T {

	if numChannels < 1 {
		numChannels = 1
	}
	if buffer < 0 {
		buffer = 0
	}

	chs := make([]chan T, numChannels)
	for i := range chs {
		chs[i] = make(chan T, buffer)
	}

	go func() {

		defer func() {
			for _, ch := range chs {
				close(ch)
			}
		}()

		for {
			val, ok := CancelOrReceive(ctx, in)
			if !ok {
				return
			}

			for _, ch := range chs {
				if !CancelOrSend(ctx, ch, val) {
					return
				}
			}
		}
	}()

	outs := make([]<-chan T, numChannels)

	for i, ch := range chs {
		outs[i] = ch
	}

	return outs
}

func Run[IN, OUT any](ctx context.Context, in <-chan IN, transform func(IN) OUT, numWorkers int, buf int) <-chan OUT {

	if numWorkers < 1 {
		panic("Must have at least 1 worker")
	}

	out := make(chan OUT, buf)
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Go(func() {
			for {
				val, ok := CancelOrReceive(ctx, in)
				if !ok {
					return
				}

				newVal := transform(val)

				if !CancelOrSend(ctx, out, newVal) {
					return
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
