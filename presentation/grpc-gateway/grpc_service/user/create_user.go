package user

import (
	"context"

	pb "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
	"github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/grpc_service/api_error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user, err := u.userCreateUseCase.CreateUser(ctx, req.Name)
	if err != nil {
		st := status.New(codes.FailedPrecondition, err.Error())
		details := api_error.FailedToCreateUser.ToGrpcErrorDetail()
		st, err := st.WithDetails(details)
		if err != nil {
			return nil, err
		}
		return nil, st.Err()
	}

	return &pb.UserResponse{
		Id:     user.UserID().Value(),
		Name:   user.UserName().Value(),
		Status: user.UserStatus().String(),
	}, nil
}
