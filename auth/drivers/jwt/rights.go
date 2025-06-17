package jwt_auth

import (
	"errors"

	"github.com/qcrg/silver_broccoli/auth"
	"github.com/qcrg/silver_broccoli/bits"
	"github.com/qcrg/silver_broccoli/database"
	"github.com/qcrg/silver_broccoli/utils"
)

type Rights struct {
	uid int64
	db  database.DB
}

func (t *Rights) ReadBalance(wallet_id int64) (bool, error) {
	privs, err := t.db.UserExtraPrivileges().GetExtraPrivileges(t.uid)
	if err != nil {
		return false, err
	}
	if utils.HasOneOf(
		privs,
		bits.UEP_BIT_GODMODE,
		bits.UEP_BIT_READ_BALANCE_FROM_ALL_WALLETS,
	) {
		return true, nil
	}

	acl, err := t.db.WalletACLs().GetACL(t.uid, wallet_id)
	if err != nil {
		return false, err
	}
	return utils.HasOneOf(acl, bits.ACL_BIT_OWNER, bits.ACL_BIT_READ_BALANCE), nil
}

func (t *Rights) FormTransactionsAsSource(wallet_id int64) (bool, error) {
	privs, err := t.db.UserExtraPrivileges().GetExtraPrivileges(t.uid)
	if err != nil {
		return false, err
	}
	if utils.HasOneOf(
		privs,
		bits.UEP_BIT_GODMODE,
		bits.UEP_BIT_FORM_TRANSACTIONS_WITH_ANY_USER_WALLETS,
	) {
		return true, nil
	}

	acl, err := t.db.WalletACLs().GetACL(t.uid, wallet_id)
	if err != nil {
		return false, err
	}

	return utils.HasOneOf(
		acl,
		bits.ACL_BIT_OWNER,
		bits.ACL_BIT_FORM_TRANSACTIONS_AS_SRC,
	), nil
}

func (t *Rights) ModifyWalletACL(wallet_id int64) (bool, error) {
	privs, err := t.db.UserExtraPrivileges().GetExtraPrivileges(t.uid)
	if err != nil {
		return false, err
	}
	if utils.HasOneOf(
		privs,
		bits.UEP_BIT_GODMODE,
		bits.UEP_BIT_MODIFY_ALL_WALLETS_ACLS,
	) {
		return true, nil
	}

	acl, err := t.db.WalletACLs().GetACL(t.uid, wallet_id)
	if err != nil {
		return false, err
	}

	return utils.HasOneOf(
		acl,
		bits.ACL_BIT_OWNER,
		bits.ACL_BIT_MODIFY_ACL,
	), nil
}

func (t *Rights) check_uep(flag bits.UserExtraPrivilegesFlags) (bool, error) {
	privs, err := t.db.UserExtraPrivileges().GetExtraPrivileges(t.uid)
	if err != nil {
		return false, err
	}
	return utils.HasOneOf(privs, bits.UEP_BIT_GODMODE, flag), nil
}

func (t *Rights) FormTransactionsWithNullWallet() (bool, error) {
	return t.check_uep(bits.UEP_BIT_FORM_TRANSACTIONS_WITH_NULL_WALLET)
}

func (t *Rights) FormNegativeBalance() (bool, error) {
	return t.check_uep(bits.UEP_BIT_FORM_NEGATIVE_BALANCE)
}

func (t *Rights) FreezingWallets() (bool, error) {
	return t.check_uep(bits.UEP_BIT_FREEZING_WALLETS)
}

func (t *Rights) FormTransactionsWithAnyUserWallets() (bool, error) {
	return t.check_uep(bits.UEP_BIT_FORM_TRANSACTIONS_WITH_ANY_USER_WALLETS)
}

func (t *Rights) ModifyUsers() (bool, error) {
	return t.check_uep(bits.UEP_BIT_MODIFY_USERS)
}

func (t *Rights) ModifyUserExtraPrivileges() (bool, error) {
	return t.check_uep(bits.UEP_BIT_MODIFY_USER_EXTRA_PRIVILEGES)
}

func NewRights(token auth.Token, db database.DB) (*Rights, error) {
	if token == nil {
		return nil, errors.New("Token is nil")
	}
	if db == nil {
		return nil, errors.New("Database is nil")
	}
	uid, err := token.GetUserId()
	if err != nil {
		return nil, err
	}
	return &Rights{uid, db}, nil
}

var _ auth.Rights = &Rights{}
