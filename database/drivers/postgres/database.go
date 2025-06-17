package postgres

import (
	"database/sql"

	"github.com/qcrg/silver_broccoli/database"
)

type BID = database.BID
type ID = database.ID

type DB struct {
	db *sql.DB
}

func (t *DB) Users() database.Users {
	return Users{t.db}
}

func (t *DB) Wallets() database.Wallets {
	return Wallets{t.db}
}

func (t *DB) WalletACLs() database.WalletACLs {
	return WalletsACLs{t.db}
}

func (t *DB) WalletTypes() database.WalletTypes {
	return WalletTypes{t.db}
}

func (t *DB) UserExtraPrivileges() database.UserExtraPrivileges {
	return UserExtraPivileges{t.db}
}

func (t *DB) Close() error {
	return t.db.Close()
}

func NewDatabase(conf Config) (*DB, error) {
	// FIXME
	if "disable" != conf.GetTLSMod() {
		log.Fatal().Msgf("Only 'disable' TLS mod is implemented")
	}

	db, err := sql.Open("postgres", conf.GetConnectionString())
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

var _ database.DB = &DB{}
