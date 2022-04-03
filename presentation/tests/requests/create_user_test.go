package requests

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/infra/ent"
	"github.com/pocket7878/go-ddd-learning/infra/ent/enttest"
)

func TestMain(m *testing.M) {
}

func TestCreateUserEndpointHave200(t *testing.T) {
	r := BuildRouter(nil)
}

func buildTestDbConnection(t *testing.T) *ent.Client {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	return client
}
