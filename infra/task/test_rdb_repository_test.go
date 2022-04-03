package task_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	taskDomain "github.com/pocket7878/go-ddd-learning/domain/task"
	"github.com/pocket7878/go-ddd-learning/infra/ent/enttest"
	entTask "github.com/pocket7878/go-ddd-learning/infra/ent/task"
	"github.com/pocket7878/go-ddd-learning/infra/task"
)

func TestSaveCreateNewTaskInDb(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	repo := task.NewTaskRdbRepository(client)
	ctx := context.Background()
	taskDue := time.Now()
	domainTask := taskDomain.NewTask("test_task", taskDue)
	err := repo.Save(ctx, domainTask)
	if err != nil {
		t.Fatalf("failed to save task: %v", err)
	}

	_, err = client.Task.Query().Where(entTask.Name(domainTask.Name())).Only(ctx)
	if err != nil {
		t.Fatalf("failed to find task: %v", err)
	}
}

func TestFindByIdCanRetrieveTaskFromDb(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	// Create data in db
	taskId := taskDomain.NewTaskID()
	_, err := client.Task.Create().
		SetID(taskId.Value()).
		SetName("test_task").
		SetTaskStatus(entTask.TaskStatus("undone")).
		SetPostponeCount(0).
		SetDueDate(time.Now()).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to save task: %v", err)
	}

	repo := task.NewTaskRdbRepository(client)
	domainTask, err := repo.FindByID(ctx, *taskId)
	if err != nil {
		t.Fatalf("failed to find task by id: %v", err)
	}

	if domainTask.TaskID().Value() != taskId.Value() {
		t.Fatalf("failed to find task by id")
	}
}
