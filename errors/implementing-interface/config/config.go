package config

import (
	"fmt"
	"os"
)

const fileHeader = "APPCONF"

type HeaderError struct {
	FaultyHeader string
}

func (e *HeaderError) Error() string {
	return fmt.Sprintf("bad header. Provide %s, expected: APPCONF", e.FaultyHeader)
}

func Load() (string, error) {
	data, err := os.ReadFile("./errors/implementing-interface/config/config.json")
	if err != nil {
		return "", err
	}
	conf := string(data)
	if conf[0:8] != fileHeader {
		return "", &HeaderError{FaultyHeader: conf[0:8]}
	}
	return conf, nil
}
