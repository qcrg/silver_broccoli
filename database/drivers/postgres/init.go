package postgres

import (
	_ "github.com/lib/pq"
	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	initiator.DefaultInitAll()
	log = initiator.GetDefaultLogger().
		With().
		Str("tag", "database").
		Str("type", "postgres").
		Logger()
}
