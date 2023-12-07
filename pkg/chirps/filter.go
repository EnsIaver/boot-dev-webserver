package chirps

import (
	"regexp"
)

func CleanChirpMessage(message string) string {
	re := regexp.MustCompile("(?i)kerfuffle|sharbert|fornax|profane")
	cleanedChirp := re.ReplaceAllString(message, "****")
	return cleanedChirp
}
