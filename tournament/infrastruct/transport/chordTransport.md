//package transport
//
//import (
//	"context"
//	"errors"
//	"net"
//	"sync"
//	"time"
//
//	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/dht"
//	pb "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_chord"
//	"google.golang.org/grpc"
//
//	log "github.com/sirupsen/logrus"
//)
//
//type BaseTransport struct {
//	config *dht.Config
//
//	sock *net.TCPListener
//
//	connPool map[string]*connect
//	pool     *sync.RWMutex
//
//	server *grpc.Server
//
//	shutdownCtx        context.Context
//	shutdownCancelFunc context.CancelFunc
//}
//
//func NewTransport(config *dht.Config) (*BaseTransport, error) {
//	listener, err := net.Listen("tcp", config.Addr)
//	if err != nil {
//		return nil, err
//	}
//	// Setup the transport
//	t := &BaseTransport{
//		config:   config,
//		sock:     listener.(*net.TCPListener),
//		connPool: make(map[string]*connect),
//		pool:     &sync.RWMutex{},
//	}
//	t.server = grpc.NewServer(config.ServerOpts...)
//
//	t.shutdownCtx, t.shutdownCancelFunc = context.WithCancel(context.Background())
//	// Done
//	return t, nil
//}
//
//func (t *BaseTransport) isShutdown() bool {
//	select {
//	case <-t.shutdownCtx.Done():
//		return true
//	default:
//		return false
//	}
//}
//
//// Gets an outbound connection to an address
//func (t *BaseTransport) getConn(addr string) (pb.ChordClient, error) {
//	t.pool.RLock()
//
//	if t.isShutdown() { // If we are shutting down, return an error
//		t.pool.RUnlock()
//		log.Error("getConn: transport is shutting down")
//		return nil, errors.New("transport is shutting down")
//	}
//
//	cc, ok := t.connPool[addr]
//	t.pool.RUnlock()
//	if ok {
//		log.Debugf("getConn: found connection for %s", addr)
//		return cc.pb_client, nil
//	}
//
//	conn, err := grpc.Dial(addr, t.config.DialOpts...) // Dial the address
//	if err != nil {
//		log.Errorf("getConn: error dialing %s: %v", addr, err)
//		return nil, err
//	}
//
//	client := pb.NewChordClient(conn)
//	cc = &connect{addr, conn, client, time.Now()}
//
//	t.pool.Lock()
//	if t.pool == nil {
//		t.pool.Unlock()
//		log.Error("getConn: transport must be initialized before calling getConn")
//		return nil, errors.New("must instantiate node before using")
//	}
//	t.connPool[addr] = cc
//	t.pool.Unlock()
//
//	return client, nil
//}
//
//func (t *BaseTransport) Start() error {
//	go t.listen() // Start listening for incoming connections
//
//	go t.reapOld()
//	return nil
//}
//
//// Stop the transport and close all connections
//func (t *BaseTransport) Stop() error {
//	t.shutdownCancelFunc()
//
//	// Close all the connections
//	t.pool.Lock()
//	t.server.Stop()
//
//	for _, conn := range t.connPool {
//		conn.Close()
//	}
//	t.connPool = nil
//	t.pool.Unlock()
//
//	return nil
//}
//
//func (t *BaseTransport) listen() {
//	t.server.Serve(t.sock)
//}
//
//// Closes old connections
//func (t *BaseTransport) reapOld() {
//	ticker := time.NewTicker(60 * time.Second)
//
//	for _ = range ticker.C {
//		if t.isShutdown() { // If we are shutting down, return
//			return
//		}
//		t.pool.Lock()
//		for host, conn := range t.connPool { // Close connections that are too old
//			if time.Since(conn.lastActive) > t.config.MaxIdle {
//				conn.Close()
//				delete(t.connPool, host)
//			}
//		}
//		t.pool.Unlock()
//	}
//}
