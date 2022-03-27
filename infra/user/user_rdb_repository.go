package user

import (
	"context"

	userDomain "github.com/pocket7878/go-ddd-learning/domain/user"
	"github.com/pocket7878/go-ddd-learning/infra/ent"
	entUser "github.com/pocket7878/go-ddd-learning/infra/ent/user"
)

type UserRdbRepository struct {
	client ent.Client
}

func NewUserRdbRepository(client ent.Client) *UserRdbRepository {
	return &UserRdbRepository{client}
}

func (r *UserRdbRepository) FindById(ctx context.Context, id userDomain.UserId) (*userDomain.User, error) {
	u, err := r.client.User.
		Query().
		Where(entUser.ID(*id.Value())).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	domainUser := userDomain.ReconstructUser(*userDomain.ReconstructUserId(u.ID), *userDomain.ReconstructUserName(u.Name), userDomain.UserStatus(u.Status))
	return domainUser, nil
}

func (r *UserRdbRepository) Save(ctx context.Context, user *userDomain.User) error {
	_, err := r.client.User.Create().
		SetID(*user.UserId().Value()).
		SetName(*user.UserName().Value()).
		SetStatus(entUser.Status(*user.UserStatus())).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
