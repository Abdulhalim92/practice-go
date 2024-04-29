package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Configuration struct {
	Server   Server `json:"server"`
	Database DB     `json:"database"`
}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DB struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	cfgBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var cfg Configuration
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
