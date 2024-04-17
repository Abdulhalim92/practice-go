package server

import "time"

// Option 3: Functional Options Pattern

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func NewServer(options ...func(server *Server)) *Server {
	srv := &Server{}

	for _, option := range options {
		option(srv)
	}

	return srv
}

func (s *Server) Start() error {
	return nil
}

func WithHost(host string) func(server *Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(server *Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(server *Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(server *Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}
