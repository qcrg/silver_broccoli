package bits

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type UserExtraPrivilegesFlags uint64

const (
	UEP_BIT_GODMODE                                 UserExtraPrivilegesFlags = 0x0001
	UEP_BIT_FORM_TRANSACTIONS_WITH_NULL_WALLET      UserExtraPrivilegesFlags = 0x0002
	UEP_BIT_FORM_NEGATIVE_BALANCE                   UserExtraPrivilegesFlags = 0x0004
	UEP_BIT_FREEZING_WALLETS                        UserExtraPrivilegesFlags = 0x0008
	UEP_BIT_FORM_TRANSACTIONS_WITH_ANY_USER_WALLETS UserExtraPrivilegesFlags = 0x0010
	UEP_BIT_MODIFY_ALL_WALLETS_ACLS                 UserExtraPrivilegesFlags = 0x0020
	UEP_BIT_READ_BALANCE_FROM_ALL_WALLETS           UserExtraPrivilegesFlags = 0x0040
	UEP_BIT_MODIFY_USERS                            UserExtraPrivilegesFlags = 0x0080
	UEP_BIT_MODIFY_USER_EXTRA_PRIVILEGES            UserExtraPrivilegesFlags = 0x0100

	UEP_BITS_COUNT int = 9
)

func (t UserExtraPrivilegesFlags) String() string {
	res := strconv.FormatUint(uint64(t), 2)
	if len(res) > UEP_BITS_COUNT {
		log.Panic().
			Msgf("Bits must be len less than or equal to %d", UEP_BITS_COUNT)
	}
	return strings.Repeat("0", UEP_BITS_COUNT-len(res)) + res
}

func ParseUserExtraPrivilegesFlags(bits string) (UserExtraPrivilegesFlags, error) {
	if len(bits) != UEP_BITS_COUNT {
		return UserExtraPrivilegesFlags(0),
			errors.New(fmt.Sprintf("The length must be %d", UEP_BITS_COUNT))
	}
	res, err := strconv.ParseUint(bits, 2, 64)
	return UserExtraPrivilegesFlags(res), err
}
