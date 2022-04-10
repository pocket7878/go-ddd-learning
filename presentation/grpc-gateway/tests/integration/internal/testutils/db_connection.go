package testutils

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/pocket7878/go-ddd-learning/infra/ent"
	"github.com/pocket7878/go-ddd-learning/infra/ent/enttest"
)

func BuildDbConnection(t *testing.T) *ent.Client {
	return enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
}
