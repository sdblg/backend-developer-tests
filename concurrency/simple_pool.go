// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"sync"
)

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
}

type Pool struct {
	sem                   *semaphore.Weighted
	ctx                   context.Context
	wg                    *sync.WaitGroup
	totalSubmissionNumber int64
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	if maxConcurrent <= 0 {
		return nil
	}

	sp := new(Pool)

	sp.sem = semaphore.NewWeighted(int64(maxConcurrent))
	sp.wg = &sync.WaitGroup{}

	sp.ctx = context.TODO()
	sp.totalSubmissionNumber = 10

	return sp
}

func (p *Pool) Submit(f func()) {

	c := int64(0)
	for ; c < p.totalSubmissionNumber; c++ {
		if err := p.sem.Acquire(p.ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v\n", err)

			return
		}
		p.wg.Add(1)

		go func(f func(), weighted *semaphore.Weighted, wg *sync.WaitGroup, c int64) {
			defer func() {
				weighted.Release(1)
				p.wg.Done()
				fmt.Printf("%vth work has been done\n", c)
			}()

			f()

		}(f, p.sem, p.wg, c)
	}

	p.wg.Wait()
}
