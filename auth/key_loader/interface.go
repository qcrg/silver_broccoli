package key_loader

import (
	"io"

	"github.com/qcrg/silver_broccoli/utils"
)

type PubKeyLoader interface {
	io.Closer
	Key(data any) (pub_key any, err error)
}

var Registry = utils.NewRegistry[PubKeyLoader]("key_loader")
