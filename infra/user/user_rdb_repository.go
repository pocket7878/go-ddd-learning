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

func (r *UserRdbRepository) FindByID(ctx context.Context, id userDomain.UserID) (*userDomain.User, error) {
	u, err := r.client.User.
		Query().
		Where(entUser.ID(id.Value())).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	domainUser := userDomain.ReconstructUser(*userDomain.ReconstructUserID(u.ID), *userDomain.ReconstructUserName(u.Name), userDomain.UserStatus(u.Status))
	return domainUser, nil
}

func (r *UserRdbRepository) Save(ctx context.Context, user *userDomain.User) error {
	_, err := r.client.User.Create().
		SetID(user.UserID().Value()).
		SetName(user.UserName().Value()).
		SetStatus(entUser.Status(*user.UserStatus())).
		OnConflict().UpdateNewValues().
		ID(ctx)

	if err != nil {
		return err
	}

	return nil
}
