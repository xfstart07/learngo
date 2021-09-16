package funcoptions

import (
	"crypto/tls"
	"time"
)

type Option func(*Server)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Server struct {
	Addr string
	Port int
	Conf Config
}

func Protocol(p string) Option {
	return func(s *Server) {
		s.Conf.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Conf.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.Conf.Maxconns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.Conf.TLS = tls
	}
}

func NewServer(addr string, port int, ops ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr: addr,
		Port: port,
		Conf: Config{
			Protocol: "tcp",
			Timeout:  30 * time.Second,
			Maxconns: 1000,
			TLS:      nil,
		},
	}
	for _, opt := range ops {
		opt(&srv)
	}

	return &srv, nil
}
