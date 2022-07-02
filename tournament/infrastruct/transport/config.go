package transport

import (
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	Addr string

	ServerOpts []grpc.ServerOption
	DialOpts   []grpc.DialOption

	Timeout time.Duration
	MaxIdle time.Duration
}

func DefaultConfig() *Config {
	n := &Config{
		Timeout: time.Second * 5,
		MaxIdle: time.Second * 10,

		ServerOpts: []grpc.ServerOption{},
		DialOpts:   make([]grpc.DialOption, 0, 5),
	}

	n.DialOpts = append(n.DialOpts,
		grpc.WithBlock(),
		grpc.WithTimeout(n.Timeout),
		grpc.FailOnNonTempDialError(true),
		grpc.WithInsecure(),
	)
	return n
}
