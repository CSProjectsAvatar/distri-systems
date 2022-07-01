package transport

import (
	"time"

	"google.golang.org/grpc"
)

type remoteCn struct {
	addr       string
	conn       *grpc.ClientConn
	lastActive time.Time
}

func (g *remoteCn) Close() {
	g.conn.Close()
}
