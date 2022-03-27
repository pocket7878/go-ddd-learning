package task

import "context"

type TaskRepository interface {
	FindById(ctx context.Context, id TaskId) (*Task, error)
	Save(ctx context.Context, task *Task) error
}
