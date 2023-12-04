package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"time"
)

func main() {
	// ErrGroup()
	ErrGroupRandom()
}

func ErrGroup() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("context 1 done, nothing to do")
			return nil
		default:
			fmt.Println("work 1 started")
			time.Sleep(1 * time.Second)
			return nil
		}
	})

	group.Go(func() error {
		fmt.Println("work 2 started")
		return fmt.Errorf("work 2 error")
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("context 3 done, nothing to do")
			return nil
		default:
			fmt.Println("work 3 started")
			time.Sleep(1 * time.Second)
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

func ErrGroupRandom() {
	group, ctx := errgroup.WithContext(context.Background())
	ch := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {

		group.Go(func() error {
			ch <- struct{}{}

			return myCallbackFunction(ctx, i)
		})

		<-ch // wait for goroutine to start to ensure they are executed in order
	}

	if err := group.Wait(); err != nil {
		log.Println(err)
	}
}

func myCallbackFunction(ctx context.Context, index int) error {
	// Print log to ensure that all works are started
	log.Printf("-- Work number %v started\n", index)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if rand.Intn(5) == 2 {
			return fmt.Errorf("work %d error", index)
		}

		// simulate work
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		fmt.Printf("Work %d completed\n", index)

		return nil
	}
}
