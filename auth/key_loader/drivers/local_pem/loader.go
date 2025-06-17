package local_pem_loader

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"reflect"

	"github.com/qcrg/silver_broccoli/auth/key_loader"
)

type KeyLoader struct {
	key any
}

func (*KeyLoader) Close() error {
	return nil
}

func (t *KeyLoader) Key(any) (any, error) {
	if t.key == nil {
		return nil, errors.New("Key is not loaded")
	}
	return t.key, nil
}

func NewKeyLoader(cfg Config) (key_loader.PubKeyLoader, error) {
	path := cfg.GetFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("Public key is not found. File: " + path)
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil,
			errors.Join(errors.New("Public key is not loaded. File: "+path), err)
	}
	log.Debug().Msgf("Key is loaded: %s", reflect.TypeOf(key).String())

	return &KeyLoader{key}, nil
}

var _ key_loader.PubKeyLoader = &KeyLoader{}
