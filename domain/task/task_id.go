package task

import (
	"github.com/google/uuid"
)

type TaskId struct {
	value string
}

func NewTaskId() *TaskId {
	uuid := uuid.New()
	return &TaskId{uuid.String()}
}

func (id *TaskId) Value() string {
	return id.value
}
