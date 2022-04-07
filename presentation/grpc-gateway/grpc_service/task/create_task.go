package task

import (
	"context"
	"time"

	pb "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
)

func (t *TaskController) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	taskDueData, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		return nil, err
	}

	task, err := t.taskCreateUseCase.CreateTask(ctx, req.Name, taskDueData)
	if err != nil {
		return nil, err
	}

	return &pb.TaskResponse{
		Id:            task.TaskID().Value(),
		Name:          task.Name(),
		Status:        task.TaskStatus().String(),
		PostponeCount: int32(task.PostponeCount()),
		DueDate:       task.DueDate().Format(time.RFC3339),
	}, nil
}
