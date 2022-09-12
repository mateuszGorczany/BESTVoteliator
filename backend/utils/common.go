package common

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ErrorNotImplemented = errors.New("Not implemented yet.")
)

var Logger *zap.Logger

func LoadConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	configDir := path.Dir(configPath)
	viper.AddConfigPath(configDir)
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	return
}

func LoadSpecificConfig(configPath string) {
	configName := path.Base(configPath)
	configDir := path.Dir(configPath)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configDir)
	viper.MergeInConfig()
}

func JSONDecodeAndValidate[T any](r io.Reader, structure *T) error {
	err := json.NewDecoder(r).Decode(&structure)
	if err != nil {
		return err
	}
	err = validator.New().Struct(*structure)
	if err != nil {
		return err
	}
	return nil
}

type ID_t string

type Stream[T any] struct {
	value []T
}

func (s Stream[T]) Collect() []T {
	return s.value
}

func (s *Stream[T]) Map(function func(element T) T) *Stream[T] {
	newList := make([]T, len(s.value))
	for i, element := range s.value {
		newList[i] = function(element)
	}
	return &Stream[T]{newList}
}

func (s *Stream[T]) Filter(filterFunction func(element T) bool) *Stream[T] {
	var newList []T
	for _, element := range s.value {
		if filterFunction(element) {
			newList = append(newList, element)
		}
	}
	return &Stream[T]{newList}
}
