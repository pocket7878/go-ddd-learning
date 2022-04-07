package grpc_service

import (
	"context"

	"github.com/pocket7878/go-ddd-learning/infra/ent"
	taskInfra "github.com/pocket7878/go-ddd-learning/infra/task"
	userInfra "github.com/pocket7878/go-ddd-learning/infra/user"
	taskGrpcService "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/grpc_service/task"
	userGrpcService "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/grpc_service/user"
	taskUseCase "github.com/pocket7878/go-ddd-learning/usecase/task"
	userUseCase "github.com/pocket7878/go-ddd-learning/usecase/user"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
)

type GrpcService struct {
	userController *userGrpcService.UserController
	taskController *taskGrpcService.TaskController
	pb.UnimplementedGoDDDLearningServiceServer
}

func NewGrpcService(client *ent.Client) *GrpcService {
	taskRepo := taskInfra.NewTaskRdbRepository(client)
	taskCreateUseCase := taskUseCase.NewTaskCreateUseCase(taskRepo)
	taskPostponeUsecCase := taskUseCase.NewTaskPostponeUseCase(taskRepo)

	userRepo := userInfra.NewUserRdbRepository(client)
	userCreateUseCase := userUseCase.NewUserCreateUseCase(userRepo)
	userDeactivateUseCase := userUseCase.NewUserDeactivateUseCase(userRepo)

	return &GrpcService{
		userController: userGrpcService.NewUserController(userCreateUseCase, userDeactivateUseCase),
		taskController: taskGrpcService.NewTaskController(taskCreateUseCase, taskPostponeUsecCase),
	}
}

func (s *GrpcService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return s.userController.CreateUser(ctx, req)
}

func (s *GrpcService) DeactivateUser(ctx context.Context, req *pb.DeactivateUserRequest) (*pb.UserResponse, error) {
	return s.userController.DeactivateUser(ctx, req)
}

func (s *GrpcService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	return s.taskController.CreateTask(ctx, req)
}

func (s *GrpcService) PostponeTask(ctx context.Context, req *pb.PostponeTaskRequest) (*emptypb.Empty, error) {
	return s.taskController.PosponeTask(ctx, req)
}
