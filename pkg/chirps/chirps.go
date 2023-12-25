package chirps

type Chirp struct {
	Id      int    `db:"id"`
	Message string `db:"message"`
}
