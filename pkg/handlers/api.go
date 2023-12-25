package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"git.standa.dev/boot-dev-webserver/pkg/chirps"
	"git.standa.dev/boot-dev-webserver/pkg/config"
	"git.standa.dev/boot-dev-webserver/pkg/database"
	"github.com/go-chi/chi/v5"
)

type BaseAPIHandler struct {
	storage database.ChirpStorage
}

func NewApiRouter(cfg *config.Config) chi.Router {
	h := BaseAPIHandler{
		storage: cfg.Storage,
	}
	r := chi.NewRouter()
	r.Get("/healthz", h.healthCheck)
	r.HandleFunc("/reset", cfg.ApiConfig.ResetMetrics)
	r.Get("/chirps", h.getChirps)
	r.Post("/chirps", h.postChirps)
	return r
}

func (h *BaseAPIHandler) healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *BaseAPIHandler) getChirps(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	var body PostChirpBody
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

type PostChirpBody struct {
	Body string
}

func (h *BaseAPIHandler) postChirps(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	var body PostChirpBody
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

	cleanedChirp := chirps.CleanChirpMessage(body.Body)
	chirp := chirps.Chirp{
		Message: cleanedChirp,
	}
	id, err := h.storage.SaveChirp(chirp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	response := ChirpResponse{
		Chirp: cleanedChirp,
		ID:    id,
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBytes)
}
