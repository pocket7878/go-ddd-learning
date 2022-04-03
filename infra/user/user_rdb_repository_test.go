package user_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	userDomain "github.com/pocket7878/go-ddd-learning/domain/user"
	"github.com/pocket7878/go-ddd-learning/infra/ent/enttest"
	entUser "github.com/pocket7878/go-ddd-learning/infra/ent/user"
	"github.com/pocket7878/go-ddd-learning/infra/user"
)

func TestSaveCreateNewUserInDb(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userName, err := userDomain.NewUserName("testuser")
	if err != nil {
		t.Fatalf("NewUserName should not return error, but got %s", err.Error())
	}
	domainUser := userDomain.NewUser(*userName)

	repo := user.NewUserRdbRepository(client)
	ctx := context.Background()
	err = repo.Save(ctx, domainUser)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	_, err = client.User.Query().Where(entUser.Name(userName.Value())).Only(ctx)
	if err != nil {
		t.Fatalf("failed to find user: %v", err)
	}
}

func TestFindByIdCanRetrieveUserFromDb(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	// Create data in db
	userId := userDomain.NewUserID()
	_, err := client.User.Create().
		SetID(userId.Value()).
		SetName("test_user").
		SetStatus("active").
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	repo := user.NewUserRdbRepository(client)
	domainUser, err := repo.FindByID(ctx, *userId)
	if err != nil {
		t.Fatalf("failed to find user by id: %v", err)
	}

	if domainUser.UserID().Value() != userId.Value() {
		t.Fatalf("failed to find user by id")
	}
}
