package main

import (
	"fmt"

	glc "github.com/HenrikFricke/go-lifecycle"
)

const (
	task1Name glc.TaskName = "task_one"
	task2Name              = "task_two"
	task3Name              = "task_three"
)

func task1() {
	fmt.Println("Task One Called")
}

func task2() {
	fmt.Println("Task Two Called")
}

func preHook() {
	fmt.Println("Pre Hook For Task Two Called")
}

func main() {
	lifecyle := glc.NewLifecyle()
	lifecyle.AddTask(task1Name, task1)
	lifecyle.AddTask(task2Name, task2)
	lifecyle.AddHook(task3Name, preHook, task2Name, glc.PRE)
	lifecyle.Execute()
}
