package jwt_auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qcrg/silver_broccoli/auth"
	"github.com/qcrg/silver_broccoli/auth/key_loader"
)

type Claims struct {
	UserId int64 `json:"uid"`
	jwt.RegisteredClaims
}

type Token struct {
	pkl    key_loader.PubKeyLoader
	claims *Claims
	token  *jwt.Token
}

func (t *Token) IsValid() bool {
	return t.token != nil && t.claims != nil
}

func (t *Token) GetUserId() (int64, error) {
	if !t.IsValid() {
		return 0, errors.New("Token is invalid")
	}
	return t.claims.UserId, nil
}

func NewToken(
	raw_token []byte,
	key_loader key_loader.PubKeyLoader,
) (*Token, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(
		string(raw_token),
		&claims,
		func(jtoken *jwt.Token) (any, error) { return key_loader.Key(jtoken) },
	)
	if err != nil {
		return nil, err
	}

	return &Token{key_loader, &claims, token}, nil
}

var _ auth.Token = &Token{}
