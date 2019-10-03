package list

import (
	"errors"
)

// Task - structure of provided task
type Task func() error

// Run - run tasks in parallel goroutines
func Run(tasks []Task, n int, maxErrors int) error {
	cntWorkers := 0
	cntFinished := 0
	cntErrors := 0
	tasksResults := make(chan error, n)
	tasksLen := len(tasks)
	for i := 0; cntFinished < tasksLen; {
		// Start new goroutine only, when:
		// 1) The number of running goroutine is less n;
		// 2) We haven't reached the end of slice;
		// 3) The amount of errors is acceptable;
		if cntWorkers < n && i < tasksLen && cntErrors < maxErrors {
			go func(task Task) {
				tasksResults <- task()
			}(tasks[i])
			i++
			cntWorkers++
		}
		select {
		case err := <-tasksResults:
			cntFinished++
			cntWorkers--
			// In case, when max errors occurred, we're waiting, when all tasks are finished and after that exiting with non-zero code
			if cntErrors >= maxErrors && cntFinished == i {
				return errors.New("too many errors occurred")
			}
			if err != nil {
				cntErrors++
			}
		// Unblock the loop's iteration
		default:
		}
	}
	return nil
}
