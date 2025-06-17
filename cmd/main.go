package main

import (
	"context"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/qcrg/silver_broccoli/api"
	"github.com/qcrg/silver_broccoli/auth"
	"github.com/qcrg/silver_broccoli/auth/key_loader"
	"github.com/qcrg/silver_broccoli/database"
	"github.com/qcrg/silver_broccoli/transport"
	"github.com/qcrg/silver_broccoli/utils"
	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/rs/zerolog/log"

	auth_drivers "github.com/qcrg/silver_broccoli/auth/drivers"
	key_loader_drivers "github.com/qcrg/silver_broccoli/auth/key_loader/drivers"
	database_drivers "github.com/qcrg/silver_broccoli/database/drivers"
	transport_drivers "github.com/qcrg/silver_broccoli/transport/drivers"
)

const (
	DATABASE_TYPE_KEY     = "DATABASE_TYPE"
	DATABASE_TYPE_DEFAULT = "postgres"

	AUTH_TYPE_KEY     = "AUTH_TYPE"
	AUTH_TYPE_DEFAULT = "file"

	TRANSPORT_TYPE_KEY     = "TRANSPORT_TYPE"
	TRANSPORT_TYPE_DEFAULT = "tcp"
)

func get_type(key string, default_ *string) string {
	if default_ == nil {
		res, present := os.LookupEnv(key)
		if !present {
			log.Fatal().Msgf("Key '%s' is not defined", key)
		}
		return res
	}
	return utils.GetEnv(key, *default_)
}

func register_drivers() {
	database_drivers.RegisterAll()
	auth_drivers.RegisterAll()
	key_loader_drivers.RegisterAll()
	transport_drivers.RegisterAll()
}

func main() {
	initiator.DefaultInitAll()
	register_drivers()

	server := api.RpcServer{Workers: api.Workers{
		Database: database.Registry.Get(get_type("DATABASE_TYPE", nil))(),
		Auth:     auth.Registry.Get(get_type("AUTH_TYPE", nil))(),
		Pkl:      key_loader.Registry.Get(get_type("AUTH_PUB_KEY_LOADER_TYPE", nil))(),
	}}

	lstn := transport.Registry.Get(get_type("TRANSPORT_TYPE", nil))()

	log.Info().Msgf("The server has started on '%s'", lstn.Addr())

	for {
		stream, err := lstn.Accept()
		if err != nil {
			panic(err)
		}
		log.Debug().
			Str("addr", stream.RemoteAddr()).
			Msg("The client has connected")

		go handle(stream, &server, context.Background())
	}
}

func handle(
	stream transport.Stream,
	server api.SilverBroccoli_Server,
	ctx context.Context,
) error {
	client := api.SilverBroccoli_ServerToClient(server)
	defer client.Release()

	conn := rpc.NewConn(
		rpc.NewStreamTransport(stream),
		&rpc.Options{BootstrapClient: capnp.Client(client)},
	)
	defer conn.Close()

	select {
	case <-ctx.Done():
		log.Debug().
			Str("addr", stream.RemoteAddr()).
			Msg("The client has disconnected")
		return nil
	case <-conn.Done():
		log.Debug().
			Str("addr", stream.RemoteAddr()).
			Msg("The client has disconnected")
		return nil
	}
}
