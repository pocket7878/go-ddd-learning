package task

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"github.com/pocket7878/go-ddd-learning/domain/task"
	pb "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
)

func (t *TaskController) PosponeTask(ctx context.Context, req *pb.PostponeTaskRequest) (*emptypb.Empty, error) {
	taskID := task.ReconstructTaskID(req.Id)
	err := t.taskPostponeUseCase.PostponeTask(ctx, *taskID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
