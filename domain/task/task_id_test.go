package task

import (
	"testing"
)

func TestNewTaskIDGenerateUniqueID(t *testing.T) {
	taskIDFirst := NewTaskID()
	taskIDSecond := NewTaskID()
	if taskIDFirst.Value() == taskIDSecond.Value() {
		t.Fatalf("NewTaskID generate same id twice")
	}
}
