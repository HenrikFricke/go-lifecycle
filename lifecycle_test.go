package golifecycle

import "testing"

func TestLifecyleExecution(t *testing.T) {
	output := ""

	var task1Name TaskName = "task_one"
	var task2Name TaskName = "task_two"
	var task3Name TaskName = "task_three"

	lifecyle := NewLifecyle()

	lifecyle.AddTask(task1Name, func() { output = output + "|task one called|" })
	lifecyle.AddTask(task2Name, func() { output = output + "|task two called|" })
	lifecyle.AddHook(task3Name, func() { output = output + "|task three called|" }, task2Name, PRE)

	lifecyle.Execute()

	if output != "|task one called||task three called||task two called|" {
		t.Error("Lifecyle execution went wrong. Order:" + output)
	}
}
