package transport_drivers

import tcp_transport "github.com/qcrg/silver_broccoli/transport/drivers/tcp"

func RegisterAll() {
	tcp_transport.RegisterTCP()
}
