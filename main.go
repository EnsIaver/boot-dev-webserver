package main

import (
	"fmt"
	"net/http"
)

const (
	hostname = "localhost"
	port = "8080"
)

func main() {
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)

	addr := fmt.Sprintf("%s:%s", hostname, port)
	http.ListenAndServe(addr, corsMux)
}
