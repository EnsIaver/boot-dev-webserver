package handlers

import (
	"net/http"

	"git.standa.dev/boot-dev-webserver/pkg/config"
	"github.com/go-chi/chi/v5"
)

func NewApiRouter(cfg *config.ApiConfig) chi.Router {
	r := chi.NewRouter()
	r.Get("/healthz", healthCheck)
	r.HandleFunc("/reset", cfg.ResetMetrics)
	return r
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
