package integrations

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/pocket7878/go-ddd-learning/infra/ent/enttest"
	entUser "github.com/pocket7878/go-ddd-learning/infra/ent/user"
	router "github.com/pocket7878/go-ddd-learning/presentation/controller/router"
)

const apiPath = "/users"

func TestCreateUserReturn201(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	r := router.BuildRouter(client)
	e := buildHttpExpect(t, r)

	e.POST(apiPath).WithJSON(map[string]interface{}{"name": "testuser"}).Expect().Status(http.StatusCreated)
}

func TestCreateUserStoreUserToDb(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	r := router.BuildRouter(client)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"name\":\"foo\"}")
	c.Request, _ = http.NewRequest("POST", apiPath, body)
	r.ServeHTTP(w, c.Request)

	usersCount, err := client.User.Query().Where(entUser.Name("foo")).Count(context.Background())
	if err != nil {
		t.Fatalf("failed to count users: %v", err)
	}
	if usersCount != 1 {
		t.Fatalf("New User not created in database")
	}
}

func buildHttpExpect(t *testing.T, router *gin.Engine) *httpexpect.Expect {
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(router),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return e
}
