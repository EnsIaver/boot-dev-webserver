package handlers

type ValidityResponse struct {
	Valid bool `json:"valid"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
