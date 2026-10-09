package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHealth(t *testing.T) {
	_ = os.Setenv("HTTP_ADDR", "127.0.0.1:0")
	s := NewServer()
	s.RegisterRoutes()

	req := httptest.NewRequest("GET", "/healthz", nil)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(`expected 200 OK, got %d`, rec.Code)
	}
}

func TestInvalidHostAddress(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic")
		}
	}()

	_ = os.Setenv("HTTP_ADDR", "")
	NewServer()
}
