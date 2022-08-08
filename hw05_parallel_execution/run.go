package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, workerCount, errorLimit int) error {
	ws := sync.WaitGroup{}
	taskCh := make(chan Task)
	errorCh := make(chan error, len(tasks))
	signalCh := make(chan struct{})

	defer func() {
		close(signalCh)
		ws.Wait()
	}()
	ws.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(taskCh, errorCh, signalCh, &ws)
	}

	for _, task := range tasks {
		taskCh <- task

		if errorLimit <= len(errorCh) {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}

func worker(taskCh chan Task, errorCh chan error, signalCh chan struct{}, ws *sync.WaitGroup) {
	for {
		select {
		case task := <-taskCh:
			if err := task(); nil != err {
				errorCh <- err
			}
		case <-signalCh:
			ws.Done()
			return
		}
	}
}
