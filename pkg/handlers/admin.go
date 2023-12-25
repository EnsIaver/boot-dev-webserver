package handlers

import (
	"git.standa.dev/boot-dev-webserver/pkg/config"
	"github.com/go-chi/chi/v5"
)

func NewAdminRouter(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Get("/metrics", cfg.ApiConfig.GetMetrics)
	return r
}
