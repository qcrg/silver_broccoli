package database

import (
	"reflect"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/stretchr/testify/mock"
)

type MockUsers struct {
	mock.Mock
}

func (t *MockUsers) Exists(id BID) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockUsers) Add(id BID) error {
	args := t.Called()
	return args.Error(0)
}

type MockWallets struct {
	mock.Mock
}

func (t *MockWallets) Exists(id BID) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockWallets) GetBalance(id BID) (BID, error) {
	args := t.Called()
	return BID(args.Int(0)), args.Error(1)
}

func (t *MockWallets) IsFrozen(id BID) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockWallets) Create(type_id ID, owner_user_id BID) (BID, error) {
	args := t.Called()
	return BID(args.Int(0)), args.Error(1)
}

type MockWalletTypes struct {
	mock.Mock
}

func (t *MockWalletTypes) Exists(type_id ID) (bool, error) {
	args := t.Called()
	return args.Bool(0), args.Error(1)
}

func (t *MockWalletTypes) Create(name string) (ID, error) {
	args := t.Called()
	return ID(args.Int(0)), args.Error(1)
}

type MockWalletACLs struct {
	mock.Mock
}

func (t *MockWalletACLs) GetACL(
	user_id BID,
	wallet_id BID,
) (bits.ACLFlags, error) {
	args := t.Called()
	arg0 := args.Get(0)
	var res bits.ACLFlags
	if reflect.TypeOf(arg0) == reflect.TypeOf(int(0)) {
		res = bits.ACLFlags(arg0.(int))
	} else {
		res = arg0.(bits.ACLFlags)
	}
	return res, args.Error(1)
}

type MockUserExtraPrivileges struct {
	mock.Mock
}

func (t *MockUserExtraPrivileges) GetExtraPrivileges(
	user_id BID,
) (bits.UserExtraPrivilegesFlags, error) {
	args := t.Called()
	arg0 := args.Get(0)
	var res bits.UserExtraPrivilegesFlags
	if reflect.TypeOf(arg0) == reflect.TypeOf(int(0)) {
		res = bits.UserExtraPrivilegesFlags(arg0.(int))
	} else {
		res = arg0.(bits.UserExtraPrivilegesFlags)
	}
	return res, args.Error(1)
}

type MockDB struct {
	mock.Mock

	Usrs   *MockUsers
	Wlts   *MockWallets
	WltTps *MockWalletTypes
	Uep    *MockUserExtraPrivileges
	Wacl   *MockWalletACLs
}

func (t *MockDB) Close() error {
	return nil
}

func (t *MockDB) Users() Users {
	return t.Usrs
}

func (t *MockDB) Wallets() Wallets {
	return t.Wlts
}

func (t *MockDB) WalletTypes() WalletTypes {
	return t.WltTps
}

func (t *MockDB) WalletACLs() WalletACLs {
	return t.Wacl
}

func (t *MockDB) UserExtraPrivileges() UserExtraPrivileges {
	return t.Uep
}

func NewMockDB() *MockDB {
	return &MockDB{
		Usrs: &MockUsers{},
		Wlts: &MockWallets{},
		Uep:  &MockUserExtraPrivileges{},
		Wacl: &MockWalletACLs{},
	}
}
