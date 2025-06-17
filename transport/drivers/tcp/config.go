package tcp_transport

import (
	"fmt"

	"github.com/qcrg/silver_broccoli/utils"
)

type Config interface {
	GetAddress() string
}

const (
	env_key_prefix = "TCP_"
	host_env_key   = env_key_prefix + "HOST"
	port_env_key   = env_key_prefix + "PORT"

	default_host = ""
	default_port = "9210"
)

type ConfigEnv struct{}

var _ Config = ConfigEnv{}

func (t ConfigEnv) GetAddress() string {
	return fmt.Sprintf("%s:%s", t.GetHost(), t.GetPort())
}

func (ConfigEnv) GetHost() string {
	return utils.GetEnv(host_env_key, default_host)
}

func (ConfigEnv) GetPort() string {
	return utils.GetEnv(port_env_key, default_port)
}
