package api

import (
	"context"
	"time"

	"github.com/qcrg/silver_broccoli/auth"
	"github.com/qcrg/silver_broccoli/auth/key_loader"
	"github.com/qcrg/silver_broccoli/database"
	"github.com/qcrg/silver_broccoli/utils/initiator"
	zlog "github.com/rs/zerolog/log"
)

func bulk(_ any) {}

type BID = database.BID
type ID = database.ID

type Workers struct {
	Database database.DB
	Auth     auth.Auth
	Pkl      key_loader.PubKeyLoader
}

type RpcServer struct {
	Workers Workers
}

func (t *RpcServer) GetBalance(
	ctx context.Context,
	call SilverBroccoli_getBalance,
) error {
	start_timestamp := time.Now()

	log := zlog.With().
		Uint64("req_id", GenRequestId()).
		Str("method", "getBalance").
		Logger()

	req := BalanceReq{}
	err := req.Deserialize(call.Args())
	if err != nil {
		log.Error().Err(err).Msg("Failed to deserialize request")
		return err
	}
	log = log.With().
		Int64("wallet_id", req.WalletId).
		Logger()

	log.Info().
		Bytes("token", []byte("*")).
		Msg("New request")

	resp, err := get_balance(req, &t.Workers, &log)
	if err != nil {
		return err
	}

	results, err := call.AllocResults()
	if err != nil {
		log.Error().Err(err).Msg("Failed to allocate results struct")
		return ErrInternal
	}

	err = resp.Serialize(results)
	if err != nil {
		log.Error().Err(err).Msg("Failed to serialize response")
		return ErrInternal
	}

	end_timestamp := time.Now()

	log.Info().
		Dur("elapsed_time_ns", end_timestamp.Sub(start_timestamp)).
		Msg("Request was performed")
	return nil
}

func (*RpcServer) GetHistory(
	ctx context.Context,
	call SilverBroccoli_getHistory,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) CreateWallet(
	ctx context.Context,
	call SilverBroccoli_createWallet,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) FreezeWallet(
	ctx context.Context,
	call SilverBroccoli_freezeWallet,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) UnfreezeWallet(
	ctx context.Context,
	call SilverBroccoli_unfreezeWallet,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) Add(
	ctx context.Context,
	call SilverBroccoli_add,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) Reduce(
	ctx context.Context,
	call SilverBroccoli_reduce,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) Transfer(
	ctx context.Context,
	call SilverBroccoli_transfer,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

func (*RpcServer) Reserve(
	ctx context.Context,
	call SilverBroccoli_reserve,
) error {
	bulk(ctx)
	bulk(call)
	return nil
}

var _ SilverBroccoli_Server = &RpcServer{}

func init() {
	initiator.DefaultInitAll()
}
