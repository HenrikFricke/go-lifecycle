package golifecycle

// Lifecycle manages tasks
type Lifecycle struct {
	mainTask *task
	tasks    map[TaskName]*task
}

// AddTask adds a task after the main process
func (l *Lifecycle) AddTask(taskName TaskName, taskFunc func(luggage interface{})) {
	newTask := &task{
		Name:    taskName,
		Process: taskFunc}

	l.tasks[taskName] = newTask
	l.mainTask.AddPostHook(newTask)
}

// AddPreHook adds a task as a pre hook
func (l *Lifecycle) AddPreHook(previousTaskName TaskName, hookName TaskName, taskFunc func(luggage interface{})) {
	newTask := &task{
		Name:    hookName,
		Process: taskFunc}

	l.tasks[hookName] = newTask
	l.tasks[previousTaskName].AddPreHook(newTask)
}

// AddPostHook adds a task as a post hook
func (l *Lifecycle) AddPostHook(previousTaskName TaskName, hookName TaskName, taskFunc func(luggage interface{})) {
	newTask := &task{
		Name:    hookName,
		Process: taskFunc}

	l.tasks[hookName] = newTask
	l.tasks[previousTaskName].AddPostHook(newTask)
}

// Execute executes the whole lifecycle
func (l *Lifecycle) Execute(luggage interface{}) {
	l.mainTask.Run(luggage)
}

// NewLifecycle returns a lifecycle
func NewLifecycle() *Lifecycle {
	mainTask := task{Process: func(l interface{}) {}}

	return &Lifecycle{
		mainTask: &mainTask,
		tasks:    make(map[TaskName]*task)}
}
