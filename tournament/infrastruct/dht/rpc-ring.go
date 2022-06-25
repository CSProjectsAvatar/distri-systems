package dht

import (
	"encoding/hex"
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"net"
	"net/rpc"
	"time"
)

type RpcRing struct {
	node *Node
	quit chan any
}

func serverName(port uint) string {
	return fmt.Sprintf("RpcRing:%d", port)
}

func (r *RpcRing) StartNode(node *Node) {
	r.node = node
	handler := rpc.NewServer()
	if err := handler.RegisterName(serverName(node.Port), r); err != nil {
		panic(err)
	}
	node.log.Info("Listening TCP...", usecases.LogArgs{"port": node.Port})
	l, e := net.Listen("tcp", fmt.Sprintf(":%d", node.Port))
	if e != nil {
		panic(e)
	}
	r.quit = make(chan any)
	go func() { // accepting connections
		for {
			select {
			case <-r.quit:
				node.log.Info(
					"Stop receiving incoming connections.",
					usecases.LogArgs{
						"Port": node.Port,
					})
				if err := l.Close(); err != nil {
					panic(err)
				}
				return
			default:
				cnx, err := l.Accept()
				if err != nil {
					panic(err)
				}
				node.log.Info(
					"RPC received.",
					usecases.LogArgs{
						"caller": cnx.RemoteAddr().String(),
						"host":   cnx.LocalAddr().String()})
				go handler.ServeConn(cnx)
			}
		}
	}()
	time.Sleep(time.Second * 2)
}

func (r *RpcRing) FindSuccRpc(id []byte, reply *RemoteNode) error {
	succ, err := r.node.FindSuccessor(id)
	if err != nil {
		return err
	}
	*reply = *succ
	return nil
}

func rpcClient(entry *RemoteNode) (*rpc.Client, error) {
	return rpc.Dial("tcp", fmt.Sprintf("%s:%d", entry.Ip, entry.Port))
}

// meth returns the service method for RPC calls.
func meth(name string, server *RemoteNode) string {
	return serverName(server.Port) + "." + name
}

func (r *RpcRing) FindSuccessor(entry *RemoteNode, id []byte) (*RemoteNode, error) {
	client, err := rpcClient(entry)
	if err != nil {
		return nil, err
	}
	var reply RemoteNode
	r.node.log.Info(
		"Calling FindSuccRpc.",
		usecases.LogArgs{
			"id":     hex.EncodeToString(id),
			"client": r.node.Port,
			"server": entry.Port})
	if err := client.Call(meth("FindSuccRpc", entry), id, &reply); err != nil {
		return nil, err
	}
	r.node.log.Info(
		"FindSuccRpc response.",
		usecases.LogArgs{
			"id":   hex.EncodeToString(id),
			"resp": reply.Addr(),
		})
	return &reply, nil
}

func (r *RpcRing) GetSuccRpc(_ *bool, reply *RemoteNode) error {
	succ, err := r.node.GetSuccessor()
	if err != nil {
		return err
	}
	*reply = *succ
	return nil
}

func (r *RpcRing) GetSuccessor(entry *RemoteNode) (*RemoteNode, error) {
	client, err := rpcClient(entry)
	if err != nil {
		return nil, err
	}
	var reply RemoteNode
	var foo bool
	if err := client.Call(meth("GetSuccRpc", entry), &foo, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

func (r *RpcRing) NotifyRpc(pred *RemoteNode, _ *int) error {
	if err := r.node.Notify(pred); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) Notify(node *RemoteNode, pred *RemoteNode) error {
	client, err := rpcClient(node)
	if err != nil {
		return err
	}
	var foo int
	if err := client.Call(meth("NotifyRpc", node), pred, &foo); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) GetPredRpc(_ *int, reply *RemoteNode) error {
	pred, err := r.node.GetPredecessor()
	if err != nil {
		return err
	}
	if pred != nil {
		*reply = *pred
	}
	return nil
}

func (r *RpcRing) GetPredecessor(node *RemoteNode) (*RemoteNode, error) {
	client, err := rpcClient(node)
	if err != nil {
		return nil, err
	}
	var reply RemoteNode
	var foo int
	r.node.log.Info(
		"Calling GetPredRpc.",
		usecases.LogArgs{
			"client": r.node.Port,
			"server": node.Port,
		})
	if err := client.Call(meth("GetPredRpc", node), &foo, &reply); err != nil {
		return nil, err
	}
	r.node.log.Info(
		"GetPredRpc response.",
		usecases.LogArgs{
			"predecessor": reply.Addr(),
			"client":      r.node.Port,
		})
	return &reply, nil
}

func (r *RpcRing) CheckNode(node *RemoteNode) error {
	_, err := rpcClient(node)
	return err
}

func (r *RpcRing) StopNode() error {
	r.quit <- 1
	time.Sleep(time.Second * 2)
	return nil
}

func (r *RpcRing) DataRpc(data []*Data, _ *int) error {
	if err := r.node.ReceiveData(data); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) SendData(data []*Data, node *RemoteNode) error {
	client, err := rpcClient(node)
	if err != nil {
		return err
	}
	var foo int
	if err := client.Call(meth("DataRpc", node), data, &foo); err != nil {
		return err
	}
	return nil
}
