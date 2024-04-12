package config

import (
	"errors"
	"os"
)

var ErrNoConfigFile = errors.New("no config file at the given path")

func Load() (string, error) {
	data, err := os.ReadFile("./errors/sentinel/config/config.json")
	if err != nil {
		return "", ErrNoConfigFile
	}
	return string(data), nil
}
