package golifecycle

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

// AddPreHook adds a task as a pre hook
func (l *Lifecyle) AddPreHook(previousTaskName TaskName, hookName TaskName, taskFunc func(luggage interface{})) {
	newTask := &task{
		Name:    hookName,
		Process: taskFunc}

	l.tasks[hookName] = newTask
	l.tasks[previousTaskName].AddPreHook(newTask)
}

// AddPostHook adds a task as a post hook
func (l *Lifecyle) AddPostHook(previousTaskName TaskName, hookName TaskName, taskFunc func(luggage interface{})) {
	newTask := &task{
		Name:    hookName,
		Process: taskFunc}

	l.tasks[hookName] = newTask
	l.tasks[previousTaskName].AddPostHook(newTask)
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
