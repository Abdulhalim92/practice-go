package server

import "time"

// Option 1: Declare a new constructor for each configuration option

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func NewServer(host string, port int) *Server {
	return &Server{
		host:    host,
		port:    port,
		timeout: 5 * time.Second,
		maxConn: 100,
	}
}

func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{
		host:    host,
		port:    port,
		timeout: timeout,
		maxConn: 100,
	}
}

func NewWithMaxConn(host string, port int, maxConn int) *Server {
	return &Server{
		host:    host,
		port:    port,
		timeout: 5 * time.Second,
		maxConn: maxConn,
	}
}

func NewWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{
		host:    host,
		port:    port,
		timeout: timeout,
		maxConn: maxConn,
	}
}

func (s *Server) Start() error {
	return nil
}
