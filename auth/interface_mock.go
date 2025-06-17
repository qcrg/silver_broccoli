package auth

import (
	"github.com/qcrg/silver_broccoli/auth/key_loader"
	"github.com/qcrg/silver_broccoli/database"
	"github.com/stretchr/testify/mock"
)

type MockToken struct {
	mock.Mock
}

func (t *MockToken) IsValid() bool {
	args := t.Called()
	return args.Bool(0)
}

func (t *MockToken) GetUserId() (int64, error) {
	args := t.Called()
	var err error
	if len(args) == 1 {
		err = nil
	} else {
		err = args.Error(1)
	}
	return int64(args.Int(0)), err
}

type MockRights struct {
	mock.Mock
}

func (t *MockRights) ReadBalance(wallet_id int64) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) FormTransactionsAsSource(wallet_id int64) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) ModifyWalletACL(wallet_id int64) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) FormTransactionsWithNullWallet() (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) FormNegativeBalance() (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) FreezingWallets() (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) FormTransactionsWithAnyUserWallets() (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) ModifyUsers() (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockRights) ModifyUserExtraPrivileges() (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

type MockAuth struct {
	Tkn  *MockToken
	Rgts *MockRights
}

func (t *MockAuth) ParseToken([]byte, key_loader.PubKeyLoader) (Token, error) {
	return t.Tkn, nil
}

func (t *MockAuth) Rights(Token, database.DB) (Rights, error) {
	return t.Rgts, nil
}

func NewMockAuth() *MockAuth {
	return &MockAuth{
		&MockToken{},
		&MockRights{},
	}
}
