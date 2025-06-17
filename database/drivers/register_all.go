package database_drivers

import (
	"github.com/qcrg/silver_broccoli/database/drivers/postgres"
)

func RegisterAll() {
	postgres.RegisterPostgres()
}
