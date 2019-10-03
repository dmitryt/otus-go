package list

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunInLimitedGoroutines(t *testing.T) {
	assert := assert.New(t)
	var wg sync.WaitGroup
	wg.Add(3)
	var calls []int
	tasks := []Task{
		func() error {
			calls[0]++
			time.Sleep(500 * time.Microsecond)
			wg.Done()
			return nil
		},
		func() error {
			calls[1]++
			time.Sleep(500 * time.Microsecond)
			wg.Done()
			return nil
		},
		func() error {
			calls[2]++
			wg.Done()
			return nil
		},
	}
	calls = make([]int, len(tasks))
	go Run(tasks, 2, 2)
	time.Sleep(50 * time.Microsecond)
	assert.Equal([]int{1,1,0}, calls, "should execute tasks in provided number of goroutines")
	wg.Wait()
	assert.Equal([]int{1,1,1}, calls, "should execute all tasks")
}

func TestRunFail(t *testing.T) {
	assert := assert.New(t)
	tasks := []Task{
		func() error {
			return nil
		},
		func() error {
			return errors.New("Some Error")
		},
		func() error {
			return errors.New("Some Error")
		},
		func() error {
			return errors.New("Some Error")
		},
	}
	assert.Equal(Run(tasks, 2, 2), errors.New("too many errors occurred"))
}

func TestWaitRunningTasks(t *testing.T) {
	assert := assert.New(t)
	var calls []int
	tasks := []Task{
		func() error {
			time.Sleep(500 * time.Microsecond)
			calls[0]++
			return nil
		},
		func() error {
			calls[1]++
			return errors.New("Some Error")
		},
		func() error {
			calls[2]++
			return nil
		},
		func() error {
			calls[3]++
			return errors.New("Some Error")
		},
		func() error {
			calls[4]++
			return errors.New("Some Error")
		},
	}
	calls = make([]int, len(tasks))
	result := Run(tasks, 2, 2)
	assert.Equal(result, errors.New("too many errors occurred"))
	assert.Equal([]int{1,1,1,1,0}, calls, "should wait for all running tasks before exiting")
}
