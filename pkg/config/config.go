package config

import (
	"fmt"
	"net/http"
	"text/template"

	"git.standa.dev/boot-dev-webserver/pkg/templates"
)

type ApiConfig struct {
	FileServerHits uint
}

func (cfg *ApiConfig) MiddlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileServerHits += 1
		fmt.Println(cfg.FileServerHits)
		next.ServeHTTP(w, r)
	})
}

func (cfg *ApiConfig) GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	t, err := template.ParseFiles("templates/admin/metrics.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	tData := templates.AdminMetricsTemplate{
		Hits: cfg.FileServerHits,
	}
	t.Execute(w, tData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *ApiConfig) ResetMetrics(w http.ResponseWriter, r *http.Request) {
	cfg.FileServerHits = 0
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
