package golifecycle

// hook describes when to run a task
type hook uint8

// Hooks
const (
	PRE hook = iota + 1
	POST
)

// Lifecyle manages tasks
type Lifecyle struct {
	mainTask *task
	tasks    map[TaskName]*task
}

// AddTask adds a task after the main process
func (l *Lifecyle) AddTask(taskName TaskName, taskFunc func(luggage interface{})) {
	newTask := &task{
		Name:    taskName,
		Process: taskFunc}

	l.tasks[taskName] = newTask
	l.mainTask.AddPostHook(newTask)
}

// AddHook adds a task as a hook
func (l *Lifecyle) AddHook(taskName TaskName, taskFunc func(luggage interface{}), previousTaskName TaskName, hook hook) {
	newTask := &task{
		Name:    taskName,
		Process: taskFunc}

	l.tasks[taskName] = newTask

	switch hook {
	case PRE:
		{
			l.tasks[previousTaskName].AddPreHook(newTask)
		}
	case POST:
		{
			l.tasks[previousTaskName].AddPostHook(newTask)
		}
	}
}

// Execute executes the whole lifecycle
func (l *Lifecyle) Execute(luggage interface{}) {
	l.mainTask.Run(luggage)
}

// NewLifecyle returns a lifecyle
func NewLifecyle() *Lifecyle {
	mainTask := task{Process: func(l interface{}) {}}

	return &Lifecyle{
		mainTask: &mainTask,
		tasks:    make(map[TaskName]*task)}
}
