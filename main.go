package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	hostname = "localhost"
	port     = "8080"

	webFilesDirectory = "static"
)

func main() {
	mux := http.NewServeMux()
	dir := http.Dir(webFilesDirectory)
	fs := http.FileServer(dir)

	cfg := apiConfig{}
	fileServerHandler := http.StripPrefix("/app", fs)
	mux.Handle("/app", cfg.middlewareMetrics(fileServerHandler))
	mux.HandleFunc("/healthz", healthCheck)
	mux.HandleFunc("/metrics", cfg.getMetrics)
	mux.HandleFunc("/reset", cfg.resetMetrics)

	corsMux := middlewareCors(mux)
	loggingMux := middlewareLogging(corsMux)

	addr := fmt.Sprintf("%s:%s", hostname, port)
	fmt.Printf("Running server on %s...\n", addr)
	err := http.ListenAndServe(addr, loggingMux)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Shutting down server...")
}
