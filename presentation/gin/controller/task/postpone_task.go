package task

import (
	"github.com/gin-gonic/gin"
	"github.com/pocket7878/go-ddd-learning/domain/task"
)

type PostponeTaskParams struct {
	ID string `json:"id"`
}

// @Schemes
// @Description postpone Task
// @Accept json
// @Produce json
// @Tags task
// @Param postponeTaskParam body PostponeTaskParams true "postpone task"
// @Success 204
// @Router /tasks/postponed [patch]
func (t *TaskController) PosponeTask(ctx *gin.Context) {
	var p PostponeTaskParams
	err := ctx.BindJSON(&p)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	taskID := task.ReconstructTaskID(p.ID)
	err = t.taskPostponeUseCase.PostponeTask(ctx, *taskID)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, gin.H{})
}
