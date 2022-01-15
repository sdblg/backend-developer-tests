package concurrency

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"sync"
)

// ErrPoolClosed is returned from AdvancedPool.Submit when the pool is closed
// before submission can be sent.
var ErrPoolClosed = errors.New("pool closed")

// AdvancedPool is a more advanced worker pool that supports cancelling the
// submission and closing the pool. All functions are safe to call from multiple
// goroutines.
type AdvancedPool interface {
	// Submit submits the given task to the pool, blocking until a slot becomes
	// available or the context is closed. The given context and its lifetime only
	// affects this function and is not the context passed to the callback. If the
	// context is closed before a slot becomes available, the context error is
	// returned. If the pool is closed before a slot becomes available,
	// ErrPoolClosed is returned. Otherwise the task is submitted to the pool and
	// no error is returned. The context passed to the callback will be closed
	// when the pool is closed.
	Submit(context.Context, func(context.Context)) error

	// Close closes the pool and waits until all submitted tasks have completed
	// before returning. If the pool is already closed, ErrPoolClosed is returned.
	// If the given context is closed before all tasks have finished, the context
	// error is returned. Otherwise, no error is returned.
	Close(context.Context) error
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.
func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	ap := new(APool)
	ap.totalSubmissionNumber = int64(maxSlots)
	ap.sem = semaphore.NewWeighted(int64(maxConcurrent))
	ap.wg = &sync.WaitGroup{}

	return ap, nil

}

type APool struct {
	isClosed bool
	Pool
}

func (A *APool) Submit(ctx context.Context, f func(context.Context)) error {
	if A.ctx == nil {
		A.ctx = ctx
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	for c := int64(0); c < A.totalSubmissionNumber; c++ {
		if err := A.sem.Acquire(A.ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v\n", err)

			return err
		}

		A.wg.Add(1)

		go func(f func(ctx context.Context), weighted *semaphore.Weighted, wg *sync.WaitGroup, c int64, ctx context.Context) {
			defer func() {
				weighted.Release(1)
				A.wg.Done()
				fmt.Printf("%vth work has been done\n\n", c)
			}()

			fmt.Printf("Starting %vth work has been started\n", c)
			f(ctx)

		}(f, A.sem, A.wg, c, ctx)
	}

	return nil
}

func (A *APool) Close(ctx context.Context) error {
	if A.isClosed {
		return ErrPoolClosed
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	A.wg.Wait()
	A.isClosed = true

	return nil
}
