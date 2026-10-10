package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/google/uuid"
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

func TestTableClick(t *testing.T) {
	cases := []struct {
		name string
		url  string
	}{
		{name: "nothing", url: "/click"},
		{name: "no offer", url: "/click?aff=1"},
		{name: "no partner", url: "/click?offer=1"},
		{name: "unknown offer", url: "/click?aff=1&offer=2"},
		{name: "unknown aff", url: "/click?aff=2&offer=1"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.url, nil)
			rec := httptest.NewRecorder()

			_ = os.Setenv("HTTP_ADDR", "127.0.0.1:0")
			s := NewServer()
			s.RegisterRoutes()

			s.router.ServeHTTP(rec, req)

			if rec.Code != http.StatusNotFound {
				t.Fatalf(`got %d`, rec.Code)
			}
		})
	}
}

func TestTableClickOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/click?aff=1&offer=1", nil)
	rec := httptest.NewRecorder()

	_ = os.Setenv("HTTP_ADDR", "127.0.0.1:0")
	s := NewServer()
	s.RegisterRoutes()

	s.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf(`got %d`, rec.Code)
	}

	loc, err := url.Parse(rec.Header().Get("Location"))
	if err != nil {
		t.Fatal(err)
	}

	if loc.Scheme != "https" || loc.Host != "google.com" || loc.Path != "" {
		t.Fatalf(`landing changed`)
	}

	tid := loc.Query().Get("tid")
	if _, err := uuid.Parse(tid); err != nil {
		t.Fatalf("tid: %v", err)
	}

	s.router.ServeHTTP(rec, req)

	loc, _ = url.Parse(rec.Header().Get("Location"))
	tid2 := loc.Query().Get("tid")

	if tid == tid2 {
		t.Fatalf(`tid2 is the same as tid`)
	}
}
