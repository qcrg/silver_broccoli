package jwt_auth

import (
	"github.com/qcrg/silver_broccoli/auth"
	"github.com/rs/zerolog/log"
)

func RegisterJWT() {
	auth.Registry.RegisterNew(
		"jwt",
		func() auth.Auth {
			auth, err := NewAuth()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create JWT Auth")
			}
			return auth
		},
	)
}
