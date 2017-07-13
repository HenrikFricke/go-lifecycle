package main

import (
	"fmt"
	"strconv"

	glc "github.com/HenrikFricke/go-lifecycle"
)

const (
	task1Name    glc.TaskName = "task_one"
	task2Name                 = "task_two"
	task3Name                 = "task_three"
	lastTaskName              = "last_task"
)

type counter struct {
	Number int
}

func (c *counter) Increment() {
	c.Number++
}

func incrementTask(luggage interface{}) {
	c := luggage.(*counter)
	c.Increment()
	fmt.Println("Task #" + strconv.Itoa(c.Number) + " called")
}

func printOut(luggage interface{}) {
	c := luggage.(*counter)
	c.Increment()
	fmt.Println("Result: " + strconv.Itoa(c.Number) + " tasks called")
}

func main() {
	c := counter{}
	lifecyle := glc.NewLifecyle()

	lifecyle.AddTask(task1Name, incrementTask)
	lifecyle.AddTask(task2Name, incrementTask)
	lifecyle.AddTask(lastTaskName, printOut)

	lifecyle.AddPreHook(task2Name, task3Name, incrementTask)

	lifecyle.Execute(&c)
}
