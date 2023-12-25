package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"git.standa.dev/boot-dev-webserver/pkg/chirps"
)

var _ ChirpStorage = &SQLiteStorage{}
var schema string = `
CREATE TABLE chirps(
	id integer primary key autoincrement,
	message text
);
`

type SQLiteStorage struct {
	db *sqlx.DB
}

func NewSQLiteStorage(db *sqlx.DB) *SQLiteStorage {
	return &SQLiteStorage{
		db: db,
	}
}

func (s *SQLiteStorage) Initialize() error {
	_, err := s.db.Exec(schema)
	return err
}

func (s *SQLiteStorage) SaveChirp(chirp chirps.Chirp) (id int64, err error) {
	query := `INSERT INTO chirps (message) VALUES (?)`
	result, err := s.db.Exec(query, chirp.Message)
	if err != nil {
		err = fmt.Errorf("can not insert into chirps: %v", err)
		return
	}

	return result.LastInsertId()
}

func (s *SQLiteStorage) GetChirps() (chirps []chirps.Chirp, err error) {
	err = s.db.Select(&chirps, "SELECT * FROM chirps")
	if err != nil {
		err = fmt.Errorf("can not select from chirps: %v", err)
		return
	}

	return
}
