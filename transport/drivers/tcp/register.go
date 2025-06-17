package tcp_transport

import (
	"github.com/qcrg/silver_broccoli/transport"
	"github.com/rs/zerolog/log"
)

func RegisterTCP() {
	transport.Registry.RegisterNew(
		"tcp",
		func() transport.Transport {
			trnsp, err := NewTransport(ConfigEnv{})
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create TCP transport")
			}
			return trnsp
		},
	)
}
