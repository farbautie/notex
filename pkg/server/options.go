package server

import (
	"net"
	"time"
)

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

func WithHost(host string) Option {
	return func(s *Server) {
		addr := s.server.Addr
		_, port, _ := net.SplitHostPort(addr)
		s.server.Addr = net.JoinHostPort(host, port)
	}
}


func WithMaxHeaderBytes(maxBytes int) Option {
	return func(s *Server) {
		s.server.MaxHeaderBytes = maxBytes
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
