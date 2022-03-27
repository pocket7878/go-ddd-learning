package task

import (
	"context"

	"github.com/pocket7878/go-ddd-learning/domain/task"
)

type TaskPostponeUseCase struct {
	taskRepository task.TaskRepository
}

func NewTaskPostponeUseCase(taskRepository task.TaskRepository) *TaskPostponeUseCase {
	return &TaskPostponeUseCase{taskRepository}
}

func (t *TaskPostponeUseCase) PostponeTask(ctx context.Context, taskId task.TaskId) error {
	task, err := t.taskRepository.FindById(ctx, taskId)
	if err != nil {
		return err
	}

	task.Postpone()

	return t.taskRepository.Save(ctx, task)
}
