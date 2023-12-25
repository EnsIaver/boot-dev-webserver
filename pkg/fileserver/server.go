package fileserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"git.standa.dev/boot-dev-webserver/pkg/config"
)

const (
	webFilesDirectory = "static"
)

func NewRouter(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	dir := http.Dir(webFilesDirectory)
	fs := http.FileServer(dir)
	fileServerHandler := http.StripPrefix("/app", fs)
	r.Handle("/*", cfg.ApiConfig.MiddlewareMetrics(fileServerHandler))
	return r
}
