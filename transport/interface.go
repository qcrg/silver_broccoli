package transport

import (
	"github.com/qcrg/silver_broccoli/utils"
)

type Stream interface {
	Close() error
	Write(src []byte) (int, error)
	Read(dst []byte) (int, error)

	RemoteAddr() string
}

type Transport interface {
	Accept() (Stream, error)
	Close() error
	Addr() string
}

var Registry = utils.NewRegistry[Transport]("transport")
