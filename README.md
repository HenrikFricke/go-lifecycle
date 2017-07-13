# go-lifecycle

## About the project

This project is inspired by the [Lifecycle](https://serverless.com/blog/writing-serverless-plugins/#lifecycle-events) of the [Serverless Framework](https://serverless.com/). The Serverless CLI handles plugins, which are able to hook before or after a process. Every plugin is also a process in the lifecycle, so other plugins can hook before or after a plugin as well. I really liked the way how Serverless build the plugin architecture and I wanted to think about how to achieve this process management in Go.

## Concept

The idea is to creeate a sequence of tasks and execute them afterwards. So we have something like a queue with tasks. The tasks will be executed with the [FIFO](https://en.wikipedia.org/wiki/FIFO_(computing_and_electronics)) method. But you can also have the special behaviour of hooks: You can add a pre or post hook to every task and also to every hook (because technically every hook is also a task).

## Example

Let's take a look to a simple example:

```go
package main

import (
	"fmt"

	glc "github.com/HenrikFricke/go-lifecycle"
)

const (
	task1Name glc.TaskName = "task_one"
	task2Name              = "task_two"
	preHook                = "pre_hook"
)

func printOut(luggage interface{}) {
	fmt.Println("I'am a task")
}

func printOutHook(luggage interface{}) {
	fmt.Println("I'am a hook")
}

func main() {
	lifecycle := glc.NewLifecycle()

	lifecycle.AddTask(task1Name, printOut)
	lifecycle.AddTask(task2Name, printOut)

	lifecycle.AddPreHook(task2Name, preHook, printOutHook)

	lifecycle.Execute(nil)
}
```

We define here two tasks and one pre hook for the second task, the output of this example is:

```bash
I'am a task
I'am a hook
I'am a task
```

You can find a more complex example in the `example` folder. You can clone this repository on your local machine and execute the example with `go run example/example.go`.

## API

You can find the documentation on [godoc](http://godoc.org/github.com/HenrikFricke/go-lifecycle).