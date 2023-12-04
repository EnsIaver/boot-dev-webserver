package config

import (
	"fmt"
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	payload := fmt.Sprintf("Hits: %d", cfg.FileServerHits)
	w.Write([]byte(payload))
}

func (cfg *ApiConfig) ResetMetrics(w http.ResponseWriter, r *http.Request) {
	cfg.FileServerHits = 0
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
