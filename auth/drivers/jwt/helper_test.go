package jwt_auth

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sign_ed25519(claims string, key ed25519.PrivateKey) []byte {
	const header = `{"alg":"EdDSA","typ":"JWT"}`
	var eheader = make([]byte, base64.RawURLEncoding.EncodedLen(len(header)))
	base64.RawURLEncoding.Encode(eheader, []byte(header))
	var eclaims = make([]byte, base64.RawURLEncoding.EncodedLen(len(claims)))
	base64.RawURLEncoding.Encode(eclaims, []byte(claims))

	signing := bytes.Join([][]byte{eheader, eclaims}, []byte("."))
	ed25519.Sign(key, signing)

	sign := make([]byte, base64.RawURLEncoding.EncodedLen(ed25519.SignatureSize))
	base64.RawURLEncoding.Encode(sign, ed25519.Sign(key, signing))

	return bytes.Join([][]byte{signing, sign}, []byte("."))
}

func TestSignEd25519(t *testing.T) {
	pkey, key, err := ed25519.GenerateKey(nil)
	require.NoError(t, err)

	raw_token := sign_ed25519(
		fmt.Sprintf(`{"exp":%d,"uid":0}`, time.Now().Add(time.Minute*10).Unix()),
		key,
	)

	token, err := jwt.Parse(
		string(raw_token),
		func(*jwt.Token) (any, error) { return pkey, nil },
	)

	{
		require.NoError(t, err)
		foo, err := token.SignedString(key)
		require.NoError(t, err)
		assert.Equal(t, foo, string(raw_token))
	}
}
