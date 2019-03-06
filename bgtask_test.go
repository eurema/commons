package arena

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	tries := 10
	interval := 500 * time.Millisecond
	incrementCounter := 0
	expectedIncrements := 6
	timerToStop := 3 * time.Second
	detectedClosing := false

	task := NewTask(func(task *Task) {
		for i := 0; i < tries; i++ {
			if task.StopRequested() {
				return
			}
			fmt.Println("Still running")
			incrementCounter++
			time.Sleep(interval)
		}
	})
	task.OnStop = func(task *Task) {
		detectedClosing = true
	}
	task.Start()
	time.Sleep(timerToStop)
	task.RequestStop()
	time.Sleep(100 * time.Millisecond) //just to wait the OnStop function do its work before the main process be killed

	assert.Equal(t, expectedIncrements, incrementCounter)
	assert.True(t, detectedClosing)
}
