package utils

import (
	"strings"

	"github.com/qcrg/silver_broccoli/utils/initiator"
	"github.com/rs/zerolog"
)

type Maker[T any] func() T

type reg_case[T any] struct {
	Make Maker[T]
}

type data_t[T any] map[string]reg_case[T]

type Registry[T any] struct {
	log  zerolog.Logger
	data data_t[T]
}

func (t *Registry[T]) RegisterNew(name string, dbm Maker[T]) {
	lname := strings.ToLower(name)
	_, exists := t.data[lname]
	if exists {
		t.log.Fatal().Msgf("Object with name '%s' was already registered", lname)
	}
	t.data[lname] = reg_case[T]{dbm}
	t.log.Debug().Msgf("New object was registered with name '%s'", lname)
}

func (t *Registry[T]) Get(name string) Maker[T] {
	lname := strings.ToLower(name)
	val, exists := t.data[lname]
	if !exists {
		t.log.Fatal().Msgf("Object with name '%s' wasn't registered", lname)
	}
	return val.Make
}

func (t *Registry[T]) GetRange() data_t[T] {
	return t.data
}

func NewRegistry[T any](type_name string) *Registry[T] {
	return &Registry[T]{
		log: initiator.GetDefaultLogger().With().
			Str("tag", "registry").
			Str("type", type_name).
			Logger(),
		data: make(data_t[T]),
	}
}
