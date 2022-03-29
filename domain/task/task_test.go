package task

import (
	"errors"
	"testing"
	"time"
)

func TestNewTaskMakeUndoneTask(t *testing.T) {
	task := NewTask("test_task", time.Now())
	if *task.TaskStatus() != Undone {
		t.Fatal("NewTask should create undone task")
	}
}

func TestNewTaskMakePostponeCountZeroTask(t *testing.T) {
	task := NewTask("test_task", time.Now())
	if task.PostponeCount() != 0 {
		t.Fatal("NewTask should create postpone count zero task")
	}
}

func TestMaxTaskPostponeIsThree(t *testing.T) {
	task := NewTask("test_task", time.Now())
	err := task.Postpone()
	if err != nil {
		t.Fatal("Postpone should be successful 1st time")
	}

	err = task.Postpone()
	if err != nil {
		t.Fatal("Postpone should be successful 2nd time")
	}

	err = task.Postpone()
	if err != nil {
		t.Fatal("Postpone should be successsful 3rd time")
	}

	err = task.Postpone()
	if err == nil {
		t.Fatal("Postpone should be failed 4th time")
	}
	if errors.Is(err, MaxPostPoneExeededError) == false {
		t.Fatal("Postpone should be failed 4th time with MaxPostPoneExeededError")
	}
}

func TestTaskDoneMakeTaskStatusToDone(t *testing.T) {
	task := NewTask("test_task", time.Now())
	task.Done()
	if *task.TaskStatus() != Done {
		t.Fatal("Done should make task status to done")
	}
}
