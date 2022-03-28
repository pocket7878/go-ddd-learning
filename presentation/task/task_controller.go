package task

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pocket7878/go-ddd-learning/domain/task"
	taskUseCase "github.com/pocket7878/go-ddd-learning/usecase/task"
)

type TaskController struct {
	taskCreateUseCase   taskUseCase.TaskCreateUseCase
	taskPostponeUseCase taskUseCase.TaskPostponeUseCase
}

func NewTaskController(taskCreateUseCase taskUseCase.TaskCreateUseCase, taskPostponeUseCase taskUseCase.TaskPostponeUseCase) *TaskController {
	return &TaskController{taskCreateUseCase, taskPostponeUseCase}
}

func (t *TaskController) CreateTask(ctx *gin.Context) {
	taskName := ctx.PostForm("name")
	taskDueDateString := ctx.PostForm("due_date")
	taskDueData, err := time.Parse(time.RFC3339, taskDueDateString)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskId, err := t.taskCreateUseCase.CreateTask(ctx, taskName, taskDueData)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, NewTaskJsonResponseBody(taskId))
}

func (t *TaskController) PosponeTask(ctx *gin.Context) {
	taskIdString := ctx.PostForm("id")
	taskId := task.ReconstructTaskId(taskIdString)
	err := t.taskPostponeUseCase.PostponeTask(ctx, *taskId)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, gin.H{})
}
