package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 0 {
		m = 0
	}
	taskCh := make(chan Task)
	var wg sync.WaitGroup
	var err error
	errmax := uint32(m)
	var errcount uint32
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				if task() != nil {
					atomic.AddUint32(&errcount, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		if atomic.LoadUint32(&errcount) >= errmax {
			err = ErrErrorsLimitExceeded
			break
		}
		taskCh <- task
	}
	close(taskCh)
	wg.Wait()
	return err
}
