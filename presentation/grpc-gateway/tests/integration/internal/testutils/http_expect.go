package testutils

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func BuildHttpExpect(t *testing.T) *httpexpect.Expect {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://localhost:8081",
		Client: &http.Client{
			Jar: httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return e
}
