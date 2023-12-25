package database

import "git.standa.dev/boot-dev-webserver/pkg/chirps"

type ChirpStorage interface {
	SaveChirp(chirp chirps.Chirp) (id int64, err error)
	GetChirps() (chirps []chirps.Chirp, err error)
}
