package postgres

import (
	"math/rand/v2"
	"testing"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/stretchr/testify/assert"
)

func TestGetExtraPrivileges(t *testing.T) {
	random_user_id := rand.Int64()
	uep := get_test_database(t).UserExtraPrivileges()

	{
		extra, err := uep.GetExtraPrivileges(random_user_id)
		assert.NoError(t, err)
		assert.Equal(t, extra, bits.UserExtraPrivilegesFlags(0))
	}
	{
		extra, err := uep.GetExtraPrivileges(0)
		assert.NoError(t, err)
		assert.Equal(t, extra, bits.UEP_BIT_GODMODE)
		assert.True(t, (extra&bits.UEP_BIT_GODMODE) != 0)
	}
	{
		extra, err := uep.GetExtraPrivileges(7)
		assert.NoError(t, err)
		assert.True(t, (extra&bits.UEP_BIT_READ_BALANCE_FROM_ALL_WALLETS) != 0)
	}
}
