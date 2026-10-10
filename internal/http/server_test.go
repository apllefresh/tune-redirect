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

func TableTestClick(t *testing.T) {
	cases := []struct {
		name string
		url  string
	}{
		{name: "nothing", url: "/click"},
		{name: "no offer", url: "/click?partner_id=1"},
		{name: "no partner", url: "/click?offer_id=1"},
		{name: "unknown offer", url: "/click?partner_id=1&offer_id=2"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.url, nil)
			rec := httptest.NewRecorder()

			s := NewServer()
			s.RegisterRoutes()

			s.router.ServeHTTP(rec, req)

			if rec.Code != http.StatusNotFound {
				t.Fatalf(`got %d`, rec.Code)
			}
		})
	}
}
