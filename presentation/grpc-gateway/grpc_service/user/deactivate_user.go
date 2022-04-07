package user

import (
	"context"

	userDomain "github.com/pocket7878/go-ddd-learning/domain/user"
	pb "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
)

func (u *UserController) DeactivateUser(ctx context.Context, req *pb.DeactivateUserRequest) (*pb.UserResponse, error) {
	userID := userDomain.ReconstructUserID(req.Id)
	user, err := u.userDeactivateUseCase.DeactivateUser(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:     user.UserID().Value(),
		Name:   user.UserName().Value(),
		Status: user.UserStatus().String(),
	}, nil
}
