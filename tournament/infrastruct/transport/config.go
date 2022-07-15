package transport

import (
	"strconv"
	"strings"

	//"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"time"

	// grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
)

type Config struct {
	Ip   string
	Port uint64

	ServerOpts []grpc.ServerOption
	DialOpts   []grpc.DialOption

	Timeout time.Duration
	MaxIdle time.Duration
}

func (c *Config) Addr() string {
	return c.Ip + ":" + strconv.Itoa(int(c.Port))
}
func DefaultCfgAddr(addr string) *Config {
	n := DefaultConfig()
	// slice addr
	n.Ip = addr[:strings.Index(addr, ":")]
	port := addr[strings.Index(addr, ":")+1:]
	// convert to uint32
	n.Port, _ = strconv.ParseUint(port, 10, 32)
	return n
}

func DefaultConfig() *Config {
	n := &Config{
		Timeout: time.Second * 8,
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
