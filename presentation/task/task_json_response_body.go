package task

import (
	"time"

	"github.com/pocket7878/go-ddd-learning/domain/task"
)

type TaskJsonResponseBody struct {
	TaskId        string    `json:"id"`
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	PostponeCount int       `json:"postpone_count"`
	DueDate       time.Time `json:"due_date"`
}

func NewTaskJsonResponseBody(task *task.Task) *TaskJsonResponseBody {
	return &TaskJsonResponseBody{
		TaskId:        task.TaskId().Value(),
		Name:          task.Name(),
		Status:        task.TaskStatus().String(),
		PostponeCount: task.PostponeCount(),
		DueDate:       task.DueDate(),
	}
}
