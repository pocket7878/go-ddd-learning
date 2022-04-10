package integration

import (
	"context"
	"net/http"
	"testing"

	entUser "github.com/pocket7878/go-ddd-learning/infra/ent/user"
	"github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/tests/integration/internal/testutils"
)

func TestCreateUserReturns200(t *testing.T) {
	cleanup, err := testutils.LaunchTestServer(t)
	defer cleanup()

	if err != nil {
		t.Fatal(err)
	}

	e := testutils.BuildHttpExpect(t)
	e.POST("/users").WithJSON(map[string]interface{}{"name": "testuser"}).Expect().Status(http.StatusOK)
}

func TestCreateUserReturnsNewlyCreatedUser(t *testing.T) {
	cleanup, err := testutils.LaunchTestServer(t)
	defer cleanup()

	if err != nil {
		t.Fatal(err)
	}

	expectedName := "testuser"
	e := testutils.BuildHttpExpect(t)
	obj := e.POST("/users").WithJSON(map[string]interface{}{"name": expectedName}).Expect().JSON().Object()
	obj.Value("id").NotNull()
	obj.Value("name").Equal(expectedName)
	obj.Value("status").Equal("inactive")
}

func TestCreateUserReturns400WithIllegalUserName(t *testing.T) {
	cleanup, err := testutils.LaunchTestServer(t)
	defer cleanup()

	if err != nil {
		t.Fatal(err)
	}

	illegalName := "@"
	e := testutils.BuildHttpExpect(t)
	e.POST("/users").WithJSON(map[string]interface{}{"name": illegalName}).Expect().Status(http.StatusBadRequest).JSON()
}

func TestCreateUserReturnsFailedToCreateUserApiError(t *testing.T) {
	cleanup, err := testutils.LaunchTestServer(t)
	defer cleanup()

	if err != nil {
		t.Fatal(err)
	}

	illegalName := "@"
	e := testutils.BuildHttpExpect(t)
	obj := e.POST("/users").WithJSON(map[string]interface{}{"name": illegalName}).Expect().JSON().Object()
	obj.Value("details").Array().Element(0).Object().Value("metadata").Object().Value("api_error_type").Equal("ApiError::FailedToCreateUser")
}

func TestCreateUserCreateNewUserInDb(t *testing.T) {
	client := testutils.BuildDbConnection(t)
	cleanup, err := testutils.LaunchTestServer(t, testutils.WithEntClient(client))
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	expectedName := "testuser"
	err = testutils.PostJSON("http://localhost:8081/users", map[string]interface{}{"name": expectedName})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	user, err := client.User.Query().Where(entUser.Name(expectedName)).Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal("New user not created")
	}
}
