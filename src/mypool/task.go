package mypool

type Task struct {
	id string
	f  func()
}

// NewTask new a task
func NewTask(f func()) *Task {
	return &Task{f: f}
}

func (t *Task) execute() {
	t.f()
}
