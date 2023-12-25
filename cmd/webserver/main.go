package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"git.standa.dev/boot-dev-webserver/pkg/config"
	"git.standa.dev/boot-dev-webserver/pkg/database"
	"git.standa.dev/boot-dev-webserver/pkg/fileserver"
	"git.standa.dev/boot-dev-webserver/pkg/handlers"
	"git.standa.dev/boot-dev-webserver/pkg/middleware"
)

const (
	hostname = "localhost"
	port     = "8080"
)

func main() {
	r := chi.NewRouter()
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("failed creating a sqlite db: %v", err)
	}

	storage := database.NewSQLiteStorage(db)
	err = storage.Initialize()
	if err != nil {
		log.Fatalf("failed initialization sqlite db: %v", err)
	}

	cfg := &config.Config{
		Storage: database.NewSQLiteStorage(db),
	}
	r.Mount("/api", handlers.NewApiRouter(cfg))
	r.Mount("/app", fileserver.NewRouter(cfg))
	r.Mount("/admin", handlers.NewAdminRouter(cfg))
	handler := middleware.Wrap(r)

	addr := fmt.Sprintf("%s:%s", hostname, port)
	log.Printf("Running server on %s...\n", addr)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("failed running the webserver: %v", err)
	}

	log.Println("Shutting down server...")
}
