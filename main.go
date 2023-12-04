package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	hostname = "localhost"
	port     = "8080"

	webFilesDirectory = "static"
)

func main() {
	r := chi.NewRouter()
	dir := http.Dir(webFilesDirectory)
	fs := http.FileServer(dir)

	cfg := apiConfig{}
	fileServerHandler := http.StripPrefix("/app", fs)
	r.Handle("/app", cfg.middlewareMetrics(fileServerHandler))
	r.Get("/healthz", healthCheck)
	r.Get("/metrics", cfg.getMetrics)
	r.HandleFunc("/reset", cfg.resetMetrics)

	corsR := middlewareCors(r)
	loggingR := middlewareLogging(corsR)

	addr := fmt.Sprintf("%s:%s", hostname, port)
	fmt.Printf("Running server on %s...\n", addr)
	err := http.ListenAndServe(addr, loggingR)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Shutting down server...")
}
