package handlers

type CleanedValidityResponse struct {
	Body string `json:"cleaned_body"`
}

type ValidityResponse struct {
	Valid bool `json:"valid"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
