package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestACLFlags_String(t *testing.T) {
	{
		acl := (ACL_BIT_MODIFY_ACL).String()
		assert.Len(t, acl, ACL_BITS_COUNT)
		assert.Equal(t, acl, "0100")
	}

	{
		acl := ACLFlags(0).String()
		assert.Len(t, acl, ACL_BITS_COUNT)
		assert.Equal(t, acl, "0000")
	}

	{
		acl := (ACL_BIT_READ_BALANCE | ACL_BIT_OWNER).String()
		assert.Len(t, acl, ACL_BITS_COUNT)
		assert.Equal(t, acl, "1001")
	}
}

func TestParseACLs(t *testing.T) {
	{
		flags, err := ParseACLs("0000")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b0000)
	}
	{
		flags, err := ParseACLs("0100")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b0100)
	}
	{
		flags, err := ParseACLs("0101")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b0101)
	}
	{
		_, err := ParseACLs("00000")
		assert.Error(t, err)
	}
	{
		_, err := ParseACLs("00")
		assert.Error(t, err)
	}
	{
		_, err := ParseACLs("")
		assert.Error(t, err)
	}
	{
		_, err := ParseACLs("0201")
		assert.Error(t, err)
	}
}

func TestUserExtraPrivilegesFlags_String(t *testing.T) {
	{
		acl := (UEP_BIT_FORM_NEGATIVE_BALANCE).String()
		assert.Len(t, acl, UEP_BITS_COUNT)
		assert.Equal(t, acl, "000000100")
	}

	{
		acl := UserExtraPrivilegesFlags(0).String()
		assert.Len(t, acl, UEP_BITS_COUNT)
		assert.Equal(t, acl, "000000000")
	}

	{
		acl := (UEP_BIT_MODIFY_USER_EXTRA_PRIVILEGES |
			UEP_BIT_FORM_TRANSACTIONS_WITH_ANY_USER_WALLETS).String()
		assert.Len(t, acl, UEP_BITS_COUNT)
		assert.Equal(t, acl, "100010000")
	}
}

func TestParseUserExtraPrivilegesFlags(t *testing.T) {
	{
		flags, err := ParseUserExtraPrivilegesFlags("000000000")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b000000000)
	}
	{
		flags, err := ParseUserExtraPrivilegesFlags("010100100")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b010100100)
	}
	{
		flags, err := ParseUserExtraPrivilegesFlags("000000100")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b000000100)
	}
	{
		flags, err := ParseUserExtraPrivilegesFlags("100000000")
		assert.NoError(t, err)
		assert.Equal(t, int(flags), 0b100000000)
	}
	{
		_, err := ParseUserExtraPrivilegesFlags("010100000000000")
		assert.Error(t, err)
	}
	{
		_, err := ParseUserExtraPrivilegesFlags("0000")
		assert.Error(t, err)
	}
	{
		_, err := ParseUserExtraPrivilegesFlags("")
		assert.Error(t, err)
	}
	{
		_, err := ParseUserExtraPrivilegesFlags("002000000")
		assert.Error(t, err)
	}
}
