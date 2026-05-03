package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

/////////////////////////////////////////////////////////////
// Err Groups
/////////////////////////////////////////////////////////////

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		fmt.Println("Goroutine 1 running")
		time.Sleep(time.Second)
		fmt.Println("Returning Error")
		return errors.New("Error Occurred")
	})

	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled")
				return nil
			default:
				fmt.Println("goroutine 2 running")
			}
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Println("Group Error: ", err)
	}
}
