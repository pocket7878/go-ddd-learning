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

func (t *TaskPostponeUseCase) PostponeTask(ctx context.Context, taskID task.TaskID) error {
	task, err := t.taskRepository.FindByID(ctx, taskID)
	if err != nil {
		return err
	}

	task.Postpone()

	return t.taskRepository.Save(ctx, task)
}
