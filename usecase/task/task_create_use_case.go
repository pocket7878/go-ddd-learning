package task

import (
	"time"

	"github.com/pocket7878/go-ddd-learning/domain/task"
)

type TaskCreateUseCase struct {
	taskRepository task.TaskRepository
}

func NewTaskCreateUseCase(taskRepository task.TaskRepository) *TaskCreateUseCase {
	return &TaskCreateUseCase{taskRepository}
}

func (t *TaskCreateUseCase) CreateTask(name string, dueDate time.Time) (*task.TaskId, error) {
	task := task.NewTask(name, dueDate)
	err := t.taskRepository.Save(task)
	if err != nil {
		return nil, err
	}

	return task.TaskId(), nil
}
