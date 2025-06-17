package jwt_auth

import (
	"errors"
	"testing"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/qcrg/silver_broccoli/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
 * The tests for errors, owner, godmode and lack of permissions are grouped
 */

func TestNewRights(t *testing.T) {
	{
		token, err := NewToken(test_data.raw_valid_token, test_data.pkl)
		require.NoError(t, err)
		rights, err := NewRights(token, database.NewMockDB())
		assert.NoError(t, err)
		assert.NotNil(t, rights)
	}
	{
		rights, err := NewRights(nil, database.NewMockDB())
		assert.Error(t, err)
		assert.Nil(t, rights)
	}
	{
		token, err := NewToken(test_data.raw_valid_token, test_data.pkl)
		require.NoError(t, err)
		rights, err := NewRights(token, nil)
		assert.Error(t, err)
		assert.Nil(t, rights)
	}
}

func make_test_rights_and_db(t *testing.T) (
	*Rights, *database.MockUserExtraPrivileges, *database.MockWalletACLs,
) {
	db := database.NewMockDB()
	token, err := NewToken(test_data.raw_valid_token, test_data.pkl)
	require.NoError(t, err)
	r, err := NewRights(token, db)
	require.NoError(t, err)
	return r, db.Uep, db.Wacl
}

func TestGodMode(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(bits.UEP_BIT_GODMODE, nil)
	wacl.On("GetACL").Return(0, nil)
	{
		has_priv, err := r.ReadBalance(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsAsSource(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.ModifyWalletACL(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsWithNullWallet()
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.FormNegativeBalance()
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.FreezingWallets()
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsWithAnyUserWallets()
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.ModifyUsers()
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.ModifyUserExtraPrivileges()
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
}

func TestNoPerms(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(0, nil)
	wacl.On("GetACL").Return(0, nil)
	{
		has_priv, err := r.ReadBalance(0)
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsAsSource(0)
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.ModifyWalletACL(0)
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsWithNullWallet()
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.FormNegativeBalance()
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.FreezingWallets()
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsWithAnyUserWallets()
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.ModifyUsers()
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
	{
		has_priv, err := r.ModifyUserExtraPrivileges()
		assert.NoError(t, err)
		assert.False(t, has_priv)
	}
}

func test_errors(t *testing.T, r *Rights, uep_must_no_error bool) {
	{
		_, err := r.ReadBalance(0)
		assert.Error(t, err)
	}
	{
		_, err := r.FormTransactionsAsSource(0)
		assert.Error(t, err)
	}
	{
		_, err := r.ModifyWalletACL(0)
		assert.Error(t, err)
	}

	{
		_, err := r.FormTransactionsWithNullWallet()
		if uep_must_no_error {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
	{
		_, err := r.FormNegativeBalance()
		if uep_must_no_error {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
	{
		_, err := r.FreezingWallets()
		if uep_must_no_error {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
	{
		_, err := r.FormTransactionsWithAnyUserWallets()
		if uep_must_no_error {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
	{
		_, err := r.ModifyUsers()
		if uep_must_no_error {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
	{
		_, err := r.ModifyUserExtraPrivileges()
		if uep_must_no_error {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestUepError(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(0, errors.New("foo"))
	wacl.On("GetACL").Return(0, nil)
	test_errors(t, r, false)
}

func TestWaclError(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(0, nil)
	wacl.On("GetACL").Return(0, errors.New("foo"))
	test_errors(t, r, true)
}

func TestUepAndWaclError(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(0, errors.New("foo"))
	wacl.On("GetACL").Return(0, errors.New("foo"))
	test_errors(t, r, false)
}

func TestOwner(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(0, nil)
	wacl.On("GetACL").Return(bits.ACL_BIT_OWNER, nil)
	{
		has_priv, err := r.ReadBalance(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.FormTransactionsAsSource(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	{
		has_priv, err := r.ModifyWalletACL(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
}

func TestReadBalance(t *testing.T) {
	// uep(READ_BALANCE) wacl(0) ret(0)
	{
		r, uep, wacl := make_test_rights_and_db(t)
		uep.On("GetExtraPrivileges").Return(
			bits.UEP_BIT_READ_BALANCE_FROM_ALL_WALLETS,
			nil,
		)
		wacl.On("GetACL").Return(0, nil)
		has_priv, err := r.ReadBalance(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
	// uep(0) wacl(READ_BALANCE) ret(true)
	{
		r, uep, wacl := make_test_rights_and_db(t)
		uep.On("GetExtraPrivileges").Return(0, nil)
		wacl.On("GetACL").Return(bits.ACL_BIT_READ_BALANCE, nil)
		has_priv, err := r.ReadBalance(0)
		assert.NoError(t, err)
		assert.True(t, has_priv)
	}
}

func TestFormTransactionsAsSource(t *testing.T) {
	{
		r, uep, wacl := make_test_rights_and_db(t)
		uep.On("GetExtraPrivileges").Return(
			bits.UEP_BIT_FORM_TRANSACTIONS_WITH_ANY_USER_WALLETS,
			nil,
		)
		wacl.On("GetACL").Return(0, nil)
		has_right, err := r.FormTransactionsAsSource(0)
		assert.NoError(t, err)
		assert.True(t, has_right)
	}
	{
		r, uep, wacl := make_test_rights_and_db(t)
		uep.On("GetExtraPrivileges").Return(0, nil)
		wacl.On("GetACL").Return(bits.ACL_BIT_FORM_TRANSACTIONS_AS_SRC, nil)
		has_right, err := r.FormTransactionsAsSource(0)
		assert.NoError(t, err)
		assert.True(t, has_right)
	}
}

func TestModifyWalletACL(t *testing.T) {
	{
		r, uep, wacl := make_test_rights_and_db(t)
		uep.On("GetExtraPrivileges").Return(
			bits.UEP_BIT_MODIFY_ALL_WALLETS_ACLS,
			nil,
		)
		wacl.On("GetACL").Return(0, nil)
		has_right, err := r.ModifyWalletACL(0)
		assert.NoError(t, err)
		assert.True(t, has_right)
	}
	{
		r, uep, wacl := make_test_rights_and_db(t)
		uep.On("GetExtraPrivileges").Return(0, nil)
		wacl.On("GetACL").Return(bits.ACL_BIT_MODIFY_ACL, nil)
		has_right, err := r.ModifyWalletACL(0)
		assert.NoError(t, err)
		assert.True(t, has_right)
	}
}

func TestFormTransactionsWithNullWallet(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(
		bits.UEP_BIT_FORM_TRANSACTIONS_WITH_NULL_WALLET,
		nil,
	)
	wacl.On("GetACL").Return(0, nil)
	has_right, err := r.FormTransactionsWithNullWallet()
	assert.NoError(t, err)
	assert.True(t, has_right)
}

func TestFormNegativeBalance(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(
		bits.UEP_BIT_FORM_NEGATIVE_BALANCE,
		nil,
	)
	wacl.On("GetACL").Return(0, nil)
	has_right, err := r.FormNegativeBalance()
	assert.NoError(t, err)
	assert.True(t, has_right)
}

func TestFreezingWallets(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(
		bits.UEP_BIT_FREEZING_WALLETS,
		nil,
	)
	wacl.On("GetACL").Return(0, nil)
	has_right, err := r.FreezingWallets()
	assert.NoError(t, err)
	assert.True(t, has_right)
}

func TestFormTransactionsWithAnyUserWallets(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(
		bits.UEP_BIT_FORM_TRANSACTIONS_WITH_ANY_USER_WALLETS,
		nil,
	)
	wacl.On("GetACL").Return(0, nil)
	has_right, err := r.FormTransactionsWithAnyUserWallets()
	assert.NoError(t, err)
	assert.True(t, has_right)
}

func TestModifyUsers(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(
		bits.UEP_BIT_MODIFY_USERS,
		nil,
	)
	wacl.On("GetACL").Return(0, nil)
	has_right, err := r.ModifyUsers()
	assert.NoError(t, err)
	assert.True(t, has_right)
}

func TestModifyUserExtraPrivileges(t *testing.T) {
	r, uep, wacl := make_test_rights_and_db(t)
	uep.On("GetExtraPrivileges").Return(
		bits.UEP_BIT_MODIFY_USER_EXTRA_PRIVILEGES,
		nil,
	)
	wacl.On("GetACL").Return(0, nil)
	has_right, err := r.ModifyUserExtraPrivileges()
	assert.NoError(t, err)
	assert.True(t, has_right)
}
