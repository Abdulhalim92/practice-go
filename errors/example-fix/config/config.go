package config

import "os"

func Load() (string, error) {
	data, err := os.ReadFile("./errors/example-1-fix/config/config.json")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
