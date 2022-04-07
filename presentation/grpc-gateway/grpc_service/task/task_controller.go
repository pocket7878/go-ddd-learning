package task

import (
	taskUseCase "github.com/pocket7878/go-ddd-learning/usecase/task"
)

type TaskController struct {
	taskCreateUseCase   *taskUseCase.TaskCreateUseCase
	taskPostponeUseCase *taskUseCase.TaskPostponeUseCase
}

func NewTaskController(taskCreateUseCase *taskUseCase.TaskCreateUseCase, taskPostponeUseCase *taskUseCase.TaskPostponeUseCase) *TaskController {
	return &TaskController{taskCreateUseCase, taskPostponeUseCase}
}
