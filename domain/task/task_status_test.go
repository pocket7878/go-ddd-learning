package task

import (
	"testing"
)

func TestUndoneTaskStatusString(t *testing.T) {
	taskStatus := TaskStatus(Undone)
	undoneString := taskStatus.String()
	if undoneString != "undone" {
		t.Fatalf("Undone.String() should be \"Undone\", but got %s", undoneString)
	}
}

func TestDoneTaskStatusString(t *testing.T) {
	taskStatus := TaskStatus(Done)
	doneString := taskStatus.String()
	if doneString != "done" {
		t.Fatalf("Undone.String() should be \"undone\", but got %s", doneString)
	}
}
