package fileserver

import (
	"net/http"

	"git.standa.dev/boot-dev-webserver/pkg/config"
	"github.com/go-chi/chi/v5"
)

const (
	webFilesDirectory = "static"
)

func NewRouter(cfg *config.ApiConfig) chi.Router {
	r := chi.NewRouter()
	dir := http.Dir(webFilesDirectory)
	fs := http.FileServer(dir)
	fileServerHandler := http.StripPrefix("/app", fs)
	r.Handle("/*", cfg.MiddlewareMetrics(fileServerHandler))
	return r
}

