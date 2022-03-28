package task

import (
	"testing"
)

func TestNewTaskIdGenerateUniqueId(t *testing.T) {
	taskIdFirst := NewTaskId()
	taskIdSecond := NewTaskId()
	if taskIdFirst.Value() == taskIdSecond.Value() {
		t.Fatalf("NewTaskId genrate same id twice")
	}
}
