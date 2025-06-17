package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"capnproto.org/go/capnp/v3/rpc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qcrg/silver_broccoli/api"
	jwt_auth "github.com/qcrg/silver_broccoli/auth/drivers/jwt"
	"github.com/qcrg/silver_broccoli/utils"
	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/rs/zerolog/log"
)

func make_token(user_id int64, token_exp time.Time) []byte {
	raw_pem, err := os.ReadFile(
		filepath.Join(utils.GetProjectDir(), "testdata/keys/ed25519_0"),
	)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(raw_pem)
	if block == nil {
		panic("pem file parse error")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil || reflect.TypeOf(key) != reflect.TypeOf(ed25519.PrivateKey{}) {
		panic(err)
	}

	claims := jwt_auth.Claims{
		UserId: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: token_exp},
		},
	}

	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)
	token_str, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}
	return []byte(token_str)
}

func main() {
	initiator.DefaultInitAll()

	user_id := flag.Int64("uid", 0, "user id")
	wallet_id := flag.Int64("wid", 0, "wallet id")
	token_exp_mul := flag.Int("", 10, "token expiration multiplier")
	flag.Parse()

	token_exp := time.Now().Add(time.Minute * time.Duration(*token_exp_mul))

	net_conn, err := net.Dial("tcp", ":9201")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	conn := rpc.NewConn(rpc.NewStreamTransport(net_conn), nil)
	defer conn.Close()

	a := api.SilverBroccoli(conn.Bootstrap(ctx))

	raw_token := make_token(*user_id, token_exp)

	future, release := a.GetBalance(
		ctx,
		func(ps api.SilverBroccoli_getBalance_Params) error {
			req := api.BalanceReq{Token: raw_token, WalletId: *wallet_id}
			req.Serialize(ps)
			return nil
		},
	)
	defer release()

	results, err := future.Struct()
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid call")
	}

	balance := results.Balance()
	if err != nil {
		panic(err)
	}

	log.Info().
		Bytes("token", []byte(raw_token)).
		Int64("user_id", *user_id).
		Int64("wallet_id", *wallet_id).
		Send()
	fmt.Printf("Wallet balance: %d\n", balance)
}
