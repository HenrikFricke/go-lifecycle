package golifecycle

import "testing"

func TestPreHookCalledInRunMethod(t *testing.T) {
	called := 0

	preTask := task{
		Process: func() { called = called + 1 }}

	mainTask := task{
		Process: func() {}}

	mainTask.AddPreHook(&preTask)
	mainTask.Run()

	if called != 1 {
		t.Error("Expected pre hook to have been called.")
	}
}

func TestPostHookCalledInRunMethod(t *testing.T) {
	called := 0

	postTask := task{
		Process: func() { called = called + 1 }}

	mainTask := task{
		Process: func() {}}

	mainTask.AddPostHook(&postTask)
	mainTask.Run()

	if called != 1 {
		t.Error("Expected post hook to have been called.")
	}
}

func TestHooksCalledInRightOrder(t *testing.T) {
	output := ""

	preOneTask := task{
		Process: func() { output = output + "|pre hook one called|" }}

	preTwoTask := task{
		Process: func() { output = output + "|pre hook two called|" }}

	postTask := task{
		Process: func() { output = output + "|post hook called|" }}

	mainTask := task{
		Process: func() { output = output + "|process called|" }}

	mainTask.AddPreHook(&preOneTask)
	mainTask.AddPostHook(&postTask)
	mainTask.AddPreHook(&preTwoTask)
	mainTask.Run()

	if output != "|pre hook one called||pre hook two called||process called||post hook called|" {
		t.Error("Hooks not called in the right order. Order:" + output)
	}
}

func TestProcessCalledInRunMethod(t *testing.T) {
	called := 0

	mainTask := task{
		Process: func() { called = called + 1 }}

	mainTask.Run()

	if called != 1 {
		t.Error("Expected process to have been called.")
	}
}
