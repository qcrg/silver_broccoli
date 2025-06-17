package postgres

import (
	"math/rand/v2"
	"testing"

	"github.com/qcrg/silver_broccoli/bits"
	"github.com/stretchr/testify/assert"
)

func TestGetACLs(t *testing.T) {
	rnd_user_id := rand.Int64()
	rnd_wallet_id := rand.Int64()
	wa := get_test_database(t).WalletACLs()

	{
		extra, err := wa.GetACL(rnd_user_id, rnd_wallet_id)
		assert.NoError(t, err)
		assert.Equal(t, extra, bits.ACLFlags(0))
	}
	{
		extra, err := wa.GetACL(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, extra, bits.ACL_BIT_OWNER)
		assert.True(t, (extra&bits.ACL_BIT_OWNER) != 0)
	}
	{
		extra, err := wa.GetACL(7, 1)
		assert.NoError(t, err)
		assert.True(t, (extra&bits.ACL_BIT_READ_BALANCE) != 0)
		assert.True(t, (extra&bits.ACL_BIT_FORM_TRANSACTIONS_AS_SRC) != 0)
		assert.False(t, (extra&bits.ACL_BIT_OWNER) != 0)
		assert.False(t, (extra&bits.ACL_BIT_MODIFY_ACL) != 0)
	}
}
