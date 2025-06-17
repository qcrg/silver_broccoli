package postgres

import (
	"database/sql"
	"errors"
)

type WalletTypes struct {
	db *sql.DB
}

func (t WalletTypes) Exists(type_id ID) (bool, error) {
	return false, errors.New("not impl")
}

func (t WalletTypes) Create(name string) (ID, error) {
	return 0, errors.New("not impl")
}
