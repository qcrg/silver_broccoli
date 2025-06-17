package api

import "github.com/rs/zerolog"

type BalanceReq struct {
	Token    []byte
	WalletId int64
}

func (t *BalanceReq) Serialize(args SilverBroccoli_getBalance_Params) error {
	err := args.SetToken(t.Token)
	if err != nil {
		return err
	}
	args.SetWalletId(t.WalletId)
	return nil
}

func (t *BalanceReq) Deserialize(args SilverBroccoli_getBalance_Params) error {
	var err error
	t.Token, err = args.Token()
	if err != nil {
		return err
	}
	t.WalletId = args.WalletId()
	return nil
}

type BalanceResp struct {
	Amount int64
}

func (t *BalanceResp) Serialize(
	results SilverBroccoli_getBalance_Results,
) error {
	results.SetBalance(t.Amount)
	return nil
}

func get_balance(
	req BalanceReq,
	t *Workers,
	log *zerolog.Logger,
) (*BalanceResp, error) {
	token, err := t.Auth.ParseToken(req.Token, t.Pkl)
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

	exist_user, err := t.Database.Users().Exists(uid)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't confirm the user's existence")
		return nil, ErrInternal
	}
	if !exist_user {
		log.Error().Msg("No user")
		return nil, ErrNoUserNoWalletNoRights
	}

	rights, err := t.Auth.Rights(token, t.Database)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get Rights object")
		return nil, ErrInternal
	}

	wallet_exists, err := t.Database.Wallets().Exists(req.WalletId)
	if err != nil {
		log.Error().Msg("Failed to check wallet existence")
		return nil, ErrInternal
	}
	if !wallet_exists {
		log.Error().Msg("No wallet")
		return nil, ErrNoUserNoWalletNoRights
	}

	has_rights, err := rights.ReadBalance(req.WalletId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check rights for wallet")
		return nil, ErrInternal
	}
	if !has_rights {
		log.Error().Msg("No rights")
		return nil, ErrNoUserNoWalletNoRights
	}

	balance, err := t.Database.Wallets().GetBalance(req.WalletId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to getBalance")
		return nil, ErrInternal
	}

	return &BalanceResp{balance}, nil
}
