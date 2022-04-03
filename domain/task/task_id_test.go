package task_test

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/domain/task"
)

func TestNewTaskIDGenerateUniqueID(t *testing.T) {
	taskIDFirst := task.NewTaskID()
	taskIDSecond := task.NewTaskID()
	if taskIDFirst.Value() == taskIDSecond.Value() {
		t.Fatalf("NewTaskID generate same id twice")
	}
}
