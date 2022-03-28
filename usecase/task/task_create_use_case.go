package task

import (
	"context"
	"time"

	"github.com/pocket7878/go-ddd-learning/domain/task"
)

type TaskCreateUseCase struct {
	taskRepository task.TaskRepository
}

func NewTaskCreateUseCase(taskRepository task.TaskRepository) *TaskCreateUseCase {
	return &TaskCreateUseCase{taskRepository}
}

func (t *TaskCreateUseCase) CreateTask(ctx context.Context, name string, dueDate time.Time) (*task.Task, error) {
	task := task.NewTask(name, dueDate)
	err := t.taskRepository.Save(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
