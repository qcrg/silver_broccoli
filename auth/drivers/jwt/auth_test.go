package jwt_auth

import (
	"crypto/ed25519"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qcrg/silver_broccoli/auth/key_loader"
	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/rs/zerolog/log"
)

type PKL struct {
	Pkey ed25519.PublicKey
}

func (PKL) Close() error {
	return nil
}

func (t *PKL) Key(any) (any, error) {
	return t.Pkey, nil
}

type TestData struct {
	raw_invalid_token []byte
	raw_valid_token   []byte
	raw_expired_token []byte
	pkey              ed25519.PublicKey
	key               ed25519.PrivateKey
	pkl               key_loader.PubKeyLoader
}

func (t *TestData) InitKeys() *TestData {
	var err error
	t.pkey, t.key, err = ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate ed25519 keys")
	}
	return t
}

func (t *TestData) InitTokens() *TestData {
	claims := Claims{
		UserId: 0,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 10)},
		},
	}

	{
		t.raw_invalid_token = sign_ed25519("invalid_claims", t.key)
	}
	{
		claims_b, err := json.Marshal(claims)
		if err != nil {
			panic(err)
		}
		t.raw_valid_token = sign_ed25519(string(claims_b), t.key)
	}
	{
		claims.ExpiresAt.Time = time.Now().Add(-time.Minute * 10)
		claims_b, err := json.Marshal(claims)
		if err != nil {
			panic(err)
		}
		t.raw_expired_token = sign_ed25519(string(claims_b), t.key)
	}

	return t
}

func (t *TestData) InitPkl() *TestData {
	t.pkl = &PKL{Pkey: t.pkey}
	return t
}

var test_data = TestData{}

func TestMain(m *testing.M) {
	test_data.InitKeys().InitTokens().InitPkl()
	initiator.InitLogger()

	os.Exit(m.Run())
}
