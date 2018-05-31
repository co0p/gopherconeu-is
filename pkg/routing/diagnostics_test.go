package routing_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/co0p/gopherconeu-is/pkg/routing"
)

func TestDiagnosticsRouter_Should_Return_StatusOK_When_HEALTHZ_is_hit(t *testing.T) {

	handler := routing.DiagnosticsRouter()
	ts := httptest.NewServer(handler)

	res, err := http.Get(ts.URL + "/healthz")
	if err != nil {
		log.Fatal(err)
	}

	expectedStatusCode := http.StatusOK
	if res.StatusCode != expectedStatusCode {
		t.Errorf("expected %v got %v", expectedStatusCode, res.StatusCode)
	}
}
