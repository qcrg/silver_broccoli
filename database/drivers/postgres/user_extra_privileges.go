package postgres

import (
	"database/sql"
	"errors"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/qcrg/silver_broccoli/database"
)

type UserExtraPivileges struct {
	db *sql.DB
}

func (t UserExtraPivileges) GetExtraPrivileges(user_id BID) (
	bits.UserExtraPrivilegesFlags, error,
) {
	row := t.db.QueryRow(
		"SELECT privileges FROM user_extra_privileges WHERE id = $1",
		user_id,
	)
	var extra string
	err := row.Scan(&extra)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return bits.UserExtraPrivilegesFlags(0), err
	}
	flags, err := bits.ParseUserExtraPrivilegesFlags(extra)
	return flags, err
}

var _ database.UserExtraPrivileges = &UserExtraPivileges{}
