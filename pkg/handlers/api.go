package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"git.standa.dev/boot-dev-webserver/pkg/chirps"
	"git.standa.dev/boot-dev-webserver/pkg/config"
	"github.com/go-chi/chi/v5"
)

func NewApiRouter(cfg *config.ApiConfig) chi.Router {
	r := chi.NewRouter()
	r.Get("/healthz", healthCheck)
	r.HandleFunc("/reset", cfg.ResetMetrics)
	r.Post("/validate_chirp", validateChirp)
	return r
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type ValidateChirpBody struct {
	Body string
}

func validateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	var body ValidateChirpBody
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	if len(body.Body) > 140 {
		response := ErrorResponse{
			Error: "Chirp too long",
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseBytes)
		return
	}

	response := CleanedValidityResponse{
		Body: chirps.CleanChirpMessage(body.Body),
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
