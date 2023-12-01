package main

import (
	"fmt"
	"net/http"
)

const (
	hostname = "localhost"
	port = "8080"

	webFilesDirectory = "static/"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(webFilesDirectory)))
	corsMux := middlewareCors(mux)

	addr := fmt.Sprintf("%s:%s", hostname, port)
	fmt.Printf("Running server on %s\n", addr)
	http.ListenAndServe(addr, corsMux)
}
