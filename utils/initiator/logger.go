package initiator

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func get_default_logger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Str("tag", "slbr").
		Logger()

	level_str := os.Getenv("LOG_LEVEL")
	if len(level_str) == 0 {
		level_str = "info"
	}
	level, err := zerolog.ParseLevel(level_str)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	zerolog.SetGlobalLevel(level)
	return logger
}

var GetDefaultLogger = sync.OnceValue(get_default_logger)
var InitLogger = sync.OnceFunc(func() { log.Logger = get_default_logger() })
