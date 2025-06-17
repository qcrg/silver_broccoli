package api

import (
	"github.com/rs/zerolog"
)

func create_wallet(
	req CreateWalletReq,
	w *Workers,
	log *zerolog.Logger,
) (*CreateWalletResp, error) {
	token, err := w.Auth.ParseToken(req.Token, w.Pkl)
	if err != nil || !token.IsValid() {
		log.Error().Err(err).Msg("Token is invalid")
		return nil, ErrInvalidToken
	}

	uid, err := token.GetUserId()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user_id from token")
		return nil, ErrInternal
	}
	*log = log.With().Int64("user_id", uid).Logger()

	exist_user, err := w.Database.Users().Exists(uid)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't confirm the user's existence")
		return nil, ErrInternal
	}
	if !exist_user {
		log.Error().Msg("No user")
		return nil, ErrNoUserNoWalletNoRights
	}

	exists, err := w.Database.WalletTypes().Exists(req.TypeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check existence of wallet type")
		return nil, ErrInternal
	}
	if !exists {
		log.Error().Err(err).Msg("Wallet type is not defined")
		return nil, ErrNoWalletType
	}

	wid, err := w.Database.Wallets().Create(req.TypeId, uid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new wallet")
		return nil, ErrInternal
	}
	*log = log.With().Int64("wallet_id", wid).Logger()

	return &CreateWalletResp{WalletId: wid}, nil
}

type CreateWalletReq struct {
	Token  []byte
	TypeId ID
}

func (t *CreateWalletReq) Deserialize(
	args SilverBroccoli_createWallet_Params,
) error {
	var err error
	t.Token, err = args.Token()
	if err != nil {
		return err
	}
	t.TypeId = args.TypeId()
	return nil
}

func (t *CreateWalletReq) Serialize(
	args SilverBroccoli_createWallet_Params,
) error {
	var err error
	err = args.SetToken(t.Token)
	if err != nil {
		return err
	}
	args.SetTypeId(t.TypeId)
	return nil
}

type CreateWalletResp struct {
	WalletId int64
}

func (t *CreateWalletResp) Deserialize(
	results SilverBroccoli_createWallet_Results,
) error {
	t.WalletId = results.WalletId()
	return nil
}

func (t *CreateWalletResp) Serialize(
	results SilverBroccoli_createWallet_Results,
) error {
	results.SetWalletId(t.WalletId)
	return nil
}
