package golifecycle

import "testing"

func TestLifecycleExecution(t *testing.T) {
	output := ""

	var task1Name TaskName = "task_one"
	var task2Name TaskName = "task_two"
	var task3Name TaskName = "task_three"

	lifecycle := NewLifecycle()

	lifecycle.AddTask(task1Name, func(luggage interface{}) { output = output + "|task one called|" })
	lifecycle.AddTask(task2Name, func(luggage interface{}) { output = output + "|task two called|" })
	lifecycle.AddPreHook(task2Name, task3Name, func(luggage interface{}) { output = output + "|task three called|" })

	lifecycle.Execute(nil)

	if output != "|task one called||task three called||task two called|" {
		t.Error("Lifecycle execution went wrong. Order:" + output)
	}
}
