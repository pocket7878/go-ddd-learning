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
	taskID        TaskID
	name          string
	taskStatus    TaskStatus
	postponeCount int
	dueDate       time.Time
}

func NewTask(name string, dueDate time.Time) *Task {
	return &Task{
		taskID:        *NewTaskID(),
		name:          name,
		taskStatus:    Undone,
		postponeCount: 0,
		dueDate:       dueDate,
	}
}

func ReconstructTask(taskID TaskID, name string, taskStatus TaskStatus, postponeCount int, dueDate time.Time) *Task {
	return &Task{
		taskID:        taskID,
		name:          name,
		taskStatus:    taskStatus,
		postponeCount: postponeCount,
		dueDate:       dueDate,
	}
}

func (t *Task) TaskID() *TaskID {
	return &t.taskID
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) TaskStatus() *TaskStatus {
	return &t.taskStatus
}

func (t *Task) PostponeCount() int {
	return t.postponeCount
}

func (t *Task) DueDate() time.Time {
	return t.dueDate
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
