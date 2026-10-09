package handler

import "net/http"

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("OK"))
	}
}

func (h *HealthCheckHandler) GetRoute() string {
	return "/healthz"
}

func (h *HealthCheckHandler) GetMethod() string {
	return http.MethodGet
}
