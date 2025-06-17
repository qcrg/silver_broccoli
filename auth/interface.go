package auth

import (
	"github.com/qcrg/silver_broccoli/auth/key_loader"
	"github.com/qcrg/silver_broccoli/database"
	"github.com/qcrg/silver_broccoli/utils"
)

type Rights interface {
	ReadBalance(wallet_id int64) (bool, error)
	FormTransactionsAsSource(wallet_id int64) (bool, error)
	ModifyWalletACL(wallet_id int64) (bool, error)

	// ReadBalanceFromAllWallets() (bool, error)
	FormTransactionsWithNullWallet() (bool, error)
	FormNegativeBalance() (bool, error)
	FreezingWallets() (bool, error)
	FormTransactionsWithAnyUserWallets() (bool, error)
	// ModifyAllWalletACLs() (bool, error)
	ModifyUsers() (bool, error)
	ModifyUserExtraPrivileges() (bool, error)
}

type Token interface {
	IsValid() bool
	GetUserId() (int64, error)
}

type Auth interface {
	ParseToken(raw_token []byte, pub_kl key_loader.PubKeyLoader) (Token, error)
	Rights(token Token, db database.DB) (Rights, error)
}

var Registry = utils.NewRegistry[Auth]("auth")
