package golifecycle

import "testing"

func TestPreHookCalledInRunMethod(t *testing.T) {
	called := 0

	preTask := task{
		Process: func(luggage interface{}) { called = called + 1 }}

	mainTask := task{
		Process: func(luggage interface{}) {}}

	mainTask.AddPreHook(&preTask)
	mainTask.Run(nil)

	if called != 1 {
		t.Error("Expected pre hook to have been called.")
	}
}

func TestPostHookCalledInRunMethod(t *testing.T) {
	called := 0

	postTask := task{
		Process: func(luggage interface{}) { called = called + 1 }}

	mainTask := task{
		Process: func(luggage interface{}) {}}

	mainTask.AddPostHook(&postTask)
	mainTask.Run(nil)

	if called != 1 {
		t.Error("Expected post hook to have been called.")
	}
}

func TestHooksCalledInRightOrder(t *testing.T) {
	output := ""

	preOneTask := task{
		Process: func(luggage interface{}) { output = output + "|pre hook one called|" }}

	preTwoTask := task{
		Process: func(luggage interface{}) { output = output + "|pre hook two called|" }}

	postTask := task{
		Process: func(luggage interface{}) { output = output + "|post hook called|" }}

	mainTask := task{
		Process: func(luggage interface{}) { output = output + "|process called|" }}

	mainTask.AddPreHook(&preOneTask)
	mainTask.AddPostHook(&postTask)
	mainTask.AddPreHook(&preTwoTask)
	mainTask.Run(nil)

	if output != "|pre hook one called||pre hook two called||process called||post hook called|" {
		t.Error("Hooks not called in the right order. Order:" + output)
	}
}

func TestProcessCalledInRunMethod(t *testing.T) {
	called := 0

	mainTask := task{
		Process: func(luggage interface{}) { called = called + 1 }}

	mainTask.Run(nil)

	if called != 1 {
		t.Error("Expected process to have been called.")
	}
}

type counter struct {
	Number int
}

func (c *counter) Increment() {
	c.Number++
}

func counterTask(luggage interface{}) {
	c := luggage.(*counter)
	c.Increment()
}

func TestLuggageTransport(t *testing.T) {
	c := counter{}

	mainTask := task{Process: counterTask}
	mainTask.AddPreHook(&task{Process: counterTask})
	mainTask.AddPostHook(&task{Process: counterTask})

	mainTask.Run(&c)

	if c.Number != 3 {
		t.Error("Luggage transport went wrong.")
	}
}
