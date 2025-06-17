package postgres

import (
	"github.com/qcrg/silver_broccoli/database"
)

func RegisterPostgres() {
	database.Registry.RegisterNew(
		"postgres",
		func() database.DB {
			db, err := NewDatabase(ConfigEnv{})
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create postgres database")
			}
			return db
		},
	)
}
