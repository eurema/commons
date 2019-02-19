package arena

// thanks https://medium.com/@matryer/stopping-goroutines-golang-1bf28799c1cb

// Task allow us to create a task in background
type Task struct {
	OnStop        func(*Task)
	task          func(*Task)
	stopChan      chan struct{}
	running       bool
	stopRequested bool
}

// NewTask creates a new task
func NewTask(task func(*Task)) *Task {
	t := new(Task)
	t.task = task
	return t
}

// Start starts the task
func (t *Task) Start() {
	t.stopChan = make(chan struct{})
	t.running = true
	go func() { // work in background
		// TODO: do setup work
		defer func() {
			if t.OnStop != nil {
				t.OnStop(t)
			}
			t.running = false
		}()
		for {
			select {
			default:
				t.task(t)
			case <-t.stopChan:
				return
			}
		}
	}()
}

// RequestStop send a stop request to the task
func (t *Task) RequestStop() {
	if !t.stopRequested {
		t.stopRequested = true
		close(t.stopChan) // tell it to stop
	}
}

// StopRequested returns true when the task was requested to stop
func (t *Task) StopRequested() bool {
	return t.stopRequested
}

// IsRunning returns true if the task it still running
func (t *Task) IsRunning() bool {
	return t.running
}
