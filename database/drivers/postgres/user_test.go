package postgres

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	var val bool
	var err error

	usr := get_test_database(t).Users()

	val, err = usr.Exists(0)
	assert.Equal(t, true, val)
	assert.NoError(t, err)

	val, err = usr.Exists(123)
	assert.Equal(t, false, val)
	assert.NoError(t, err)

	val, err = usr.Exists(5)
	assert.Equal(t, false, val)
	assert.NoError(t, err)

	val, err = usr.Exists(9)
	assert.Equal(t, true, val)
	assert.NoError(t, err)
}

func TestAddUser(t *testing.T) {
	var err error
	var user_id = rand.Int64()

	usr := get_test_database(t).Users()

	err = usr.Add(user_id)
	assert.NoError(t, err)

	err = usr.Add(user_id)
	assert.Error(t, err)
}
