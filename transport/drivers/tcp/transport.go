package tcp_transport

import (
	"net"

	"github.com/qcrg/silver_broccoli/transport"
)

type Stream struct {
	conn net.Conn
}

func (t *Stream) Close() error {
	return t.conn.Close()
}

func (t *Stream) Read(dst []byte) (int, error) {
	return t.conn.Read(dst)
}

func (t *Stream) Write(src []byte) (int, error) {
	return t.conn.Write(src)
}

func (t *Stream) RemoteAddr() string {
	return t.conn.RemoteAddr().String()
}

type Transport struct {
	lstn net.Listener
}

func (t *Transport) Accept() (transport.Stream, error) {
	conn, err := t.lstn.Accept()
	return &Stream{conn}, err
}

func (t *Transport) Close() error {
	return t.lstn.Close()
}

func (t *Transport) Addr() string {
	return t.lstn.Addr().String()
}

func NewTransport(cfg Config) (*Transport, error) {
	lstn, err := net.Listen("tcp", cfg.GetAddress())
	if err != nil {
		return nil, err
	}
	return &Transport{lstn}, nil
}

var _ transport.Transport = &Transport{}
