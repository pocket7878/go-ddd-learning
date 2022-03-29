package task

import "context"

type TaskRepository interface {
	FindByID(ctx context.Context, id TaskID) (*Task, error)
	Save(ctx context.Context, task *Task) error
}
