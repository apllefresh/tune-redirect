package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	s := NewServer("127.0.0.1:0")
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

	NewServer("")
}
