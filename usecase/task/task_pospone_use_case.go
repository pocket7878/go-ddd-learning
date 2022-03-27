package task

import (
	"github.com/pocket7878/go-ddd-learning/domain/task"
)

type TaskPostponeUseCase struct {
	taskRepository task.TaskRepository
}

func NewTaskPostponeUseCase(taskRepository task.TaskRepository) *TaskPostponeUseCase {
	return &TaskPostponeUseCase{taskRepository}
}

func (t *TaskPostponeUseCase) PostponeTask(taskId task.TaskId) error {
	task, err := t.taskRepository.FindById(taskId)
	if err != nil {
		return err
	}

	task.Postpone()

	return t.taskRepository.Save(task)
}
