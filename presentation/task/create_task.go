package task

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreateTaskParams struct {
	Name    string `json:"name"`
	DueDate string `json:"due_date"`
}

// @Schemes
// @Description create Task
// @Accept json
// @Produce json
// @Tags task
// @Param createTaskParam body CreateTaskParams true "create task"
// @Success 200 {object} task.TaskJsonResponseBody
// @Router /tasks [post]
func (t *TaskController) CreateTask(ctx *gin.Context) {
	var p CreateTaskParams
	err := ctx.BindJSON(&p)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	taskDueData, err := time.Parse(time.RFC3339, p.DueDate)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskID, err := t.taskCreateUseCase.CreateTask(ctx, p.Name, taskDueData)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, NewTaskJsonResponseBody(taskID))
}
