package handlers

import (
	"git.standa.dev/boot-dev-webserver/pkg/config"
	"github.com/go-chi/chi/v5"
)

func NewAdminRouter(cfg *config.ApiConfig) chi.Router {
	r := chi.NewRouter()
	r.Get("/metrics", cfg.GetMetrics)
	return r
}
