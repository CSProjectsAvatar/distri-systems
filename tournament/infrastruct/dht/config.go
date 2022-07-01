package dht

import (
	"time"

	"google.golang.org/grpc"
)

type Node interface {
	// trans GrpcTransport
}

type KV struct {
}

type Config struct {
	Id   string
	Addr string

	ServerOpts []grpc.ServerOption
	DialOpts   []grpc.DialOption

	Timeout time.Duration
	MaxIdle time.Duration
}
