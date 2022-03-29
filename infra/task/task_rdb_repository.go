package task

import (
	"context"

	taskDomain "github.com/pocket7878/go-ddd-learning/domain/task"
	"github.com/pocket7878/go-ddd-learning/infra/ent"
	entTask "github.com/pocket7878/go-ddd-learning/infra/ent/task"
)

type TaskRdbRepository struct {
	client ent.Client
}

func NewTaskRdbRepository(client ent.Client) *TaskRdbRepository {
	return &TaskRdbRepository{client}
}

func (r *TaskRdbRepository) FindByID(ctx context.Context, id taskDomain.TaskID) (*taskDomain.Task, error) {
	t, err := r.client.Task.Query().Where(entTask.ID(id.Value())).Only(ctx)
	if err != nil {
		return nil, err
	}

	domainTask := taskDomain.ReconstructTask(*taskDomain.ReconstructTaskID(t.ID), t.Name, taskDomain.TaskStatus(t.TaskStatus), t.PostponeCount, t.DueDate)

	return domainTask, nil
}

func (r *TaskRdbRepository) Save(ctx context.Context, task *taskDomain.Task) error {
	_, err := r.client.Task.Create().
		SetID(task.TaskID().Value()).
		SetName(task.Name()).
		SetTaskStatus(entTask.TaskStatus(*task.TaskStatus())).
		SetPostponeCount(task.PostponeCount()).
		SetDueDate(task.DueDate()).
		OnConflict().UpdateNewValues().
		ID(ctx)

	if err != nil {
		return err
	}

	return nil
}
