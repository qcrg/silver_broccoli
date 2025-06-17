package jwt_auth

import (
	"github.com/qcrg/silver_broccoli/auth"
	"github.com/qcrg/silver_broccoli/auth/key_loader"
	"github.com/qcrg/silver_broccoli/database"
)

type Auth struct{}

func (*Auth) ParseToken(
	raw_token []byte,
	pub_kl key_loader.PubKeyLoader,
) (auth.Token, error) {
	return NewToken(raw_token, pub_kl)
}

func (*Auth) Rights(token auth.Token, db database.DB) (auth.Rights, error) {
	return NewRights(token, db)
}

func NewAuth() (*Auth, error) {
	return &Auth{}, nil
}

var _ auth.Auth = &Auth{}
