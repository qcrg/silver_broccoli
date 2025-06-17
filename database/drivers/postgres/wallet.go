package postgres

import (
	"database/sql"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/qcrg/silver_broccoli/database"
)

type Wallets struct {
	db *sql.DB
}

func (t Wallets) Exists(id BID) (bool, error) {
	row := t.db.QueryRow("SELECT COUNT(*) FROM wallets WHERE id = $1", id)
	var count int = 0
	err := row.Scan(&count)
	return count > 0, err
}

func (t Wallets) GetBalance(id BID) (BID, error) {
	row := t.db.QueryRow("SELECT amount FROM wallets WHERE id = $1", id)
	var amount BID = 0
	err := row.Scan(&amount)
	return amount, err
}

func (t Wallets) IsFrozen(id BID) (bool, error) {
	row := t.db.QueryRow("SELECT frozen FROM wallets WHERE id = $1", id)
	var is_frozen bool
	err := row.Scan(&is_frozen)
	return is_frozen, err
}

func (t Wallets) Create(type_id ID, user_id BID) (BID, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	row := tx.QueryRow(
		"INSERT INTO wallets(type_id) VALUES ($1) RETURNING id", type_id)
	var wallet_id BID
	err = row.Scan(&wallet_id)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(
		"INSERT INTO wallet_acls (user_id, wallet_id, access_rights) VALUES ($1, $2, $3)",
		user_id,
		wallet_id,
		bits.ACL_BIT_OWNER.String(),
	)
	if err != nil {
		return 0, err
	}

	return wallet_id, tx.Commit()
}

var _ database.Wallets = Wallets{}
