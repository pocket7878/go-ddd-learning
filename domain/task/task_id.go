package task

import (
	"github.com/google/uuid"
)

type TaskID struct {
	value string
}

func NewTaskID() *TaskID {
	uuid := uuid.New()
	return &TaskID{uuid.String()}
}

func ReconstructTaskID(value string) *TaskID {
	return &TaskID{value}
}

func (id *TaskID) Value() string {
	return id.value
}
