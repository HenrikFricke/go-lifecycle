package golifecycle

// TaskName describes the name of a task
// it has to be unique within a Lifecyle
type TaskName string

// task handles a process with hooks
type task struct {
	Name      TaskName
	Process   func()
	preHooks  []*task
	postHooks []*task
}

// Run runs the task with hooks
func (t task) Run() {
	// run pre hooks
	for _, task := range t.preHooks {
		task.Run()
	}

	// run task itself
	t.Process()

	// run post hooks
	for _, task := range t.postHooks {
		task.Run()
	}
}

// AddPreHook adds a hook to be called PRE the task
func (t *task) AddPreHook(hook *task) {
	t.preHooks = append(t.preHooks, hook)
}

// AddPostHook adds a hook to be called POST the task
func (t *task) AddPostHook(hook *task) {
	t.postHooks = append(t.postHooks, hook)
}
