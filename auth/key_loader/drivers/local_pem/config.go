package local_pem_loader

import "os"

type Config interface {
	GetFilePath() string
}

type ConfigEnv struct{}

const (
	key_prefix    = "AUTH_PUB_KEY_LOADER_"
	file_path_key = key_prefix + "FILE_PATH"
)

var _ Config = &ConfigEnv{}

func (ConfigEnv) GetFilePath() string {
	res, present := os.LookupEnv(file_path_key)
	if !present {
		log.Fatal().Msgf("File path with key '%s' not defined", file_path_key)
	}
	return res
}
