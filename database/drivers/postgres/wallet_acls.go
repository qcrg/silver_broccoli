package postgres

import (
	"database/sql"
	"errors"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/qcrg/silver_broccoli/database"
)

type WalletsACLs struct {
	db *sql.DB
}

func (t WalletsACLs) GetACL(
	user_id int64,
	wallet_id int64,
) (bits.ACLFlags, error) {
	row := t.db.QueryRow(
		"SELECT access_rights FROM wallet_acls WHERE user_id = $1 AND wallet_id = $2",
		user_id,
		wallet_id,
	)
	var acl string
	err := row.Scan(&acl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return bits.ACLFlags(0), err
	}
	res, err := bits.ParseACLs(acl)
	return res, err
}

var _ database.WalletACLs = WalletsACLs{}
