package local_pem_loader

import (
	"github.com/qcrg/silver_broccoli/auth/key_loader"
)

func RegisterLocalPEM() {
	key_loader.Registry.RegisterNew(
		"local_pem",
		func() key_loader.PubKeyLoader {
			pkl, err := NewKeyLoader(ConfigEnv{})
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create new local_pem key loader")
			}
			return pkl
		},
	)
}
