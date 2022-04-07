package user

import (
	"context"

	pb "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
)

func (u *UserController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user, err := u.userCreateUseCase.CreateUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:     user.UserID().Value(),
		Name:   user.UserName().Value(),
		Status: user.UserStatus().String(),
	}, nil
}
