package bits

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type ACLFlags uint64

const (
	ACL_BIT_OWNER                     ACLFlags = 0x0001
	ACL_BIT_FORM_TRANSACTIONS_AS_SRC  ACLFlags = 0x0002
	ACL_BIT_MODIFY_ACL ACLFlags = 0x0004
	ACL_BIT_READ_BALANCE              ACLFlags = 0x0008

	ACL_BITS_COUNT int = 4
)

func (t ACLFlags) String() string {
	res := strconv.FormatUint(uint64(t), 2)
	if len(res) > ACL_BITS_COUNT {
		log.Panic().
			Msgf("Bits must be len less than or equal to %d", ACL_BITS_COUNT)
	}
	return strings.Repeat("0", ACL_BITS_COUNT-len(res)) + res
}

func ParseACLs(bits string) (ACLFlags, error) {
	if len(bits) != ACL_BITS_COUNT {
		return ACLFlags(0),
			errors.New(fmt.Sprintf("The length must be %d", ACL_BITS_COUNT))
	}
	res, err := strconv.ParseUint(bits, 2, 64)
	return ACLFlags(res), err
}
