package config

import (
	"errors"
	"os"
)

const fileHeader = "APPCONF"

func Load() (string, error) {
	data, err := os.ReadFile("./errors/example-fix-2/config/config.json")
	if err != nil {
		return "", err
	}

	conf := string(data)
	if conf[0:8] != fileHeader {
		return "", errors.New("the config file do not begin by accepted header")
		//return "", fmt.Errorf("the config file do not begin by accepted header")
	}

	return conf, nil

}
