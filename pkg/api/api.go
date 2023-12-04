package api

import (
	"net/http"

	"git.standa.dev/boot-dev-webserver/pkg/config"
	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg *config.ApiConfig) chi.Router {
	r := chi.NewRouter()
	r.Get("/healthz", healthCheck)
	r.Get("/metrics", cfg.GetMetrics)
	r.HandleFunc("/reset", cfg.ResetMetrics)
	return r
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

