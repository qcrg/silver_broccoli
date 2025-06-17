package database

import (
	"io"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/qcrg/silver_broccoli/utils"
)

type BID = int64
type ID = int32

type Users interface {
	Exists(id BID) (bool, error)
	Add(id BID) error
}

type Wallets interface {
	Exists(id BID) (bool, error)
	GetBalance(id BID) (int64, error)
	IsFrozen(id BID) (bool, error)
	Create(type_id ID, owner_user_id BID) (int64, error)
}

type WalletTypes interface {
	Exists(id ID) (bool, error)
}

type WalletACLs interface {
	GetACL(user_id BID, wallet_id BID) (bits.ACLFlags, error)
}

type UserExtraPrivileges interface {
	GetExtraPrivileges(user_id BID) (bits.UserExtraPrivilegesFlags, error)
}

type DB interface {
	io.Closer
	Users() Users
	Wallets() Wallets
	WalletACLs() WalletACLs
	WalletTypes() WalletTypes
	UserExtraPrivileges() UserExtraPrivileges
}

var Registry = utils.NewRegistry[DB]("database")
