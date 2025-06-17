package postgres

import (
	"database/sql"

	"github.com/qcrg/silver_broccoli/database"
)

type Users struct {
	db *sql.DB
}

func (t Users) Exists(id BID) (bool, error) {
	row := t.db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", id)
	var count int = 0
	err := row.Scan(&count)
	return count > 0, err
}

func (t Users) Add(id BID) error {
	_, err := t.db.Exec("INSERT INTO users VALUES ($1)", id)
	return err
}

var _ database.Users = Users{}
