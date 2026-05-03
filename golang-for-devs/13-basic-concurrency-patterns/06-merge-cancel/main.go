package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Merge Cancel (New Or Channel)
/////////////////////////////////////////////////////////////

func ExternalContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func main() {
	ctx1, cancel1 := ExternalContext()
	ctx2, _ := ExternalContext()
	ctx3, _ := ExternalContext()

	parent := context.Background()
	mergedContext, cancel := MergeCancel(parent, ctx1, ctx2, ctx3)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Go(func() {
		for {
			select {
			case <-mergedContext.Done():
				return
			default:
				fmt.Println("Running")
			}
		}
	})

	time.AfterFunc(2*time.Second, func() {
		cancel1()
	})

	wg.Wait()
}

func MergeCancel(parent context.Context, ctxs ...context.Context) (context.Context, context.CancelFunc) {

	mergedContext, cancel := context.WithCancelCause(parent)

	stops := make([]func() bool, len(ctxs))

	for i, ctx := range ctxs {
		stop := context.AfterFunc(ctx, func() {
			cancel(context.Cause(ctx))
		})

		stops[i] = stop
	}

	context.AfterFunc(mergedContext, func() {
		for _, stop := range stops {
			stop()
		}
	})

	return mergedContext, func() {
		cancel(nil)
	}
}
