package task_test

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/domain/task"
)

func TestUndoneTaskStatusString(t *testing.T) {
	taskStatus := task.TaskStatus(task.Undone)
	undoneString := taskStatus.String()
	if undoneString != "undone" {
		t.Fatalf("Undone.String() should be \"Undone\", but got %s", undoneString)
	}
}

func TestDoneTaskStatusString(t *testing.T) {
	taskStatus := task.TaskStatus(task.Done)
	doneString := taskStatus.String()
	if doneString != "done" {
		t.Fatalf("Undone.String() should be \"undone\", but got %s", doneString)
	}
}
