package main

import (
	"fmt"
	"log"
	"net/http"

	"git.standa.dev/boot-dev-webserver/pkg/handlers"
	"git.standa.dev/boot-dev-webserver/pkg/config"
	"git.standa.dev/boot-dev-webserver/pkg/fileserver"
	"git.standa.dev/boot-dev-webserver/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	hostname = "localhost"
	port     = "8080"
)

func main() {
	r := chi.NewRouter()
	cfg := &config.ApiConfig{}
	r.Mount("/api", handlers.NewApiRouter(cfg))
	r.Mount("/app", fileserver.NewRouter(cfg))
	r.Mount("/admin", handlers.NewAdminRouter(cfg))
	handler := middleware.Wrap(r)

	addr := fmt.Sprintf("%s:%s", hostname, port)
	fmt.Printf("Running server on %s...\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Shutting down server...")
}
