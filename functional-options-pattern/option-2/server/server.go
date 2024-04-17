package server

import "time"

// Option 2: Use a custom Config struct

type Server struct {
	cfg Config
}

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	return nil
}
