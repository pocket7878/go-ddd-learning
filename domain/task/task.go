package task

import (
	"errors"
	"time"
)

const (
	maxPostPoneCount = 3
)

var (
	MaxPostPoneExeededError = errors.New("MaxPostPoneExeededError")
)

type Task struct {
	taskId        TaskId
	name          string
	taskStatus    TaskStatus
	postponeCount int
	dueDate       time.Time
}

func NewTask(name string, dueDate time.Time) *Task {
	return &Task{
		taskId:        *NewTaskId(),
		name:          name,
		taskStatus:    Undone,
		postponeCount: 0,
		dueDate:       dueDate,
	}
}

func (t *Task) TaskId() *TaskId {
	return &t.taskId
}

func (task *Task) Postpone() error {
	err := task.validatePostPoneCount()
	if err != nil {
		return err
	}

	task.dueDate = task.dueDate.AddDate(0, 0, 1)
	task.postponeCount += 1

	return nil
}

func (task *Task) Done() {
	task.taskStatus = Done
}

func (task *Task) validatePostPoneCount() error {
	if task.postponeCount >= maxPostPoneCount {
		return MaxPostPoneExeededError
	}

	return nil
}
