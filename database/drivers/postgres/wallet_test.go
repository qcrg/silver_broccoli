package postgres

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	var err error
	var amount int64
	var rand_wallet_id = rand.Int64()
	wlt := get_test_database(t).Wallets()

	_, err = wlt.GetBalance(rand_wallet_id)
	assert.Error(t, err)

	amount, err = wlt.GetBalance(1)
	assert.Equal(t, amount, int64(100))
	assert.NoError(t, err)

	amount, err = wlt.GetBalance(2)
	assert.Equal(t, amount, int64(200))
	assert.NoError(t, err)

	amount, err = wlt.GetBalance(5)
	assert.Equal(t, amount, int64(-100))
	assert.NoError(t, err)
}

func TestIsFrozen(t *testing.T) {
	var err error
	var frozen bool
	rand_wallet_id := rand.Int64()
	wlt := get_test_database(t).Wallets()

	_, err = wlt.IsFrozen(rand_wallet_id)
	assert.Error(t, err)

	frozen, err = wlt.IsFrozen(0)
	assert.Equal(t, frozen, false)
	assert.NoError(t, err)

	frozen, err = wlt.IsFrozen(1)
	assert.Equal(t, frozen, false)
	assert.NoError(t, err)

	frozen, err = wlt.IsFrozen(2)
	assert.Equal(t, frozen, true)
	assert.NoError(t, err)

	frozen, err = wlt.IsFrozen(5)
	assert.Equal(t, frozen, true)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	var wallet_id int64
	var err error
	rand_type_id := rand.Int32()
	wlt := get_test_database(t).Wallets()
	log.Warn().Int32("int", rand_type_id).Send()

	wallet_id, err = wlt.Create(rand_type_id, 0)
	assert.Error(t, err)

	wallet_id, err = wlt.Create(0, 0)
	assert.NoError(t, err)
	assert.NotEqual(t, wallet_id, 0)

	wallet_id, err = wlt.Create(1, 0)
	assert.NoError(t, err)
}
