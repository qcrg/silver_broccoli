package local_pem_loader

import (
	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	initiator.DefaultInitAll()

	log = initiator.GetDefaultLogger().With().
		Str("tag", "key_loaders").
		Str("type", "local_pem").
		Logger()
}
