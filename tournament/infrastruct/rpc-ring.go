package infrastruct

import (
	"encoding/hex"
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"net"
	"net/rpc"
	"time"
)

type RpcRing struct {
	server *chord.Node
	quit   chan any
	client *chord.RemoteNode
	log    domain.Logger
}

func serverName(port uint) string {
	return fmt.Sprintf("RpcRing:%d", port)
}

func (r *RpcRing) StartNode(node *chord.Node) {
	r.server = node
	handler := rpc.NewServer()
	if err := handler.RegisterName(serverName(node.Port), r); err != nil {
		panic(err)
	}
	node.Log.Info("Listening TCP...", domain.LogArgs{"port": node.Port})
	l, e := net.Listen("tcp", fmt.Sprintf(":%d", node.Port))
	if e != nil {
		panic(e)
	}
	r.quit = make(chan any)
	go func() { // accepting connections
		for {
			select {
			case <-r.quit:
				node.Log.Info(
					"Stop receiving incoming connections.",
					domain.LogArgs{
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
				node.Log.Trace(
					"RPC received.",
					domain.LogArgs{
						"caller": cnx.RemoteAddr().String(),
						"host":   cnx.LocalAddr().String()})
				go handler.ServeConn(cnx)
			}
		}
	}()
	time.Sleep(time.Second * 2)
}

func (r *RpcRing) FindSuccRpc(id []byte, reply *chord.RemoteNode) error {
	succ, err := r.server.FindSuccessor(id)
	if err != nil {
		return err
	}
	*reply = *succ
	return nil
}

func rpcClient(entry *chord.RemoteNode) (*rpc.Client, error) {
	return rpc.Dial("tcp", fmt.Sprintf("%s:%d", entry.Ip, entry.Port))
}

// meth returns the service method for RPC calls.
func meth(name string, server *chord.RemoteNode) string {
	return serverName(server.Port) + "." + name
}

func (r *RpcRing) FindSuccessor(entry *chord.RemoteNode, id []byte) (*chord.RemoteNode, error) {
	client, err := rpcClient(entry)
	if err != nil {
		return nil, err
	}
	var reply chord.RemoteNode
	r.log.Trace(
		"Calling FindSuccRpc.",
		domain.LogArgs{
			"id":     hex.EncodeToString(id),
			"client": r.client.Port,
			"server": entry.Port})
	if err := client.Call(meth("FindSuccRpc", entry), id, &reply); err != nil {
		return nil, err
	}
	r.log.Trace(
		"FindSuccRpc response.",
		domain.LogArgs{
			"id":   hex.EncodeToString(id),
			"resp": reply.Addr(),
		})
	return &reply, nil
}

func (r *RpcRing) GetSuccRpc(_ *bool, reply *chord.RemoteNode) error {
	succ := r.server.GetSuccessor()
	*reply = *succ
	return nil
}

func (r *RpcRing) GetSuccessor(entry *chord.RemoteNode) (*chord.RemoteNode, error) {
	client, err := rpcClient(entry)
	if err != nil {
		return nil, err
	}
	var reply chord.RemoteNode
	var foo bool
	if err := client.Call(meth("GetSuccRpc", entry), &foo, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

func (r *RpcRing) NotifyRpc(pred *chord.RemoteNode, _ *int) error {
	if err := r.server.Notify(pred); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) Notify(node *chord.RemoteNode, pred *chord.RemoteNode) error {
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

func (r *RpcRing) GetPredRpc(_ *int, reply *chord.RemoteNode) error {
	pred := r.server.GetPredecessor()
	if pred != nil {
		*reply = *pred
	}
	return nil
}

func (r *RpcRing) GetPredecessor(node *chord.RemoteNode) (*chord.RemoteNode, error) {
	client, err := rpcClient(node)
	if err != nil {
		return nil, err
	}
	var reply chord.RemoteNode
	var foo int
	r.log.Trace(
		"Calling GetPredRpc.",
		domain.LogArgs{
			"client": r.client.Port,
			"server": node.Port,
		})
	if err := client.Call(meth("GetPredRpc", node), &foo, &reply); err != nil {
		return nil, err
	}
	r.log.Trace(
		"GetPredRpc response.",
		domain.LogArgs{
			"predecessor": reply.Addr(),
			"client":      r.client.Port,
		})
	return &reply, nil
}

func (r *RpcRing) CheckNode(node *chord.RemoteNode) error {
	_, err := rpcClient(node)
	return err
}

func (r *RpcRing) StopNode() error {
	close(r.quit)
	time.Sleep(time.Second * 2)
	return nil
}

func (r *RpcRing) DataRpc(data []*chord.Data, _ *int) error {
	if err := r.server.ReceiveData(data); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) SendData(data []*chord.Data, node *chord.RemoteNode) error {
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

func (r *RpcRing) SetValue(node *chord.RemoteNode, key []byte, value string) error {
	client, err := rpcClient(node)
	if err != nil {
		return err
	}
	var foo int
	args := chord.Data{
		Key: key, Value: value,
	}
	if err := client.Call(meth("SetValueRpc", node), &args, &foo); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) SetValueRpc(data *chord.Data, _ *int) error {
	if err := r.server.SetValue(data.Key, data.Value); err != nil {
		return err
	}
	return nil
}

func (r *RpcRing) GetValue(node *chord.RemoteNode, key []byte) (string, error) {
	client, err := rpcClient(node)
	if err != nil {
		return "", err
	}
	var value string
	if err := client.Call(meth("GetValueRpc", node), key, &value); err != nil {
		return "", err
	}
	return value, nil
}

func (r *RpcRing) GetValueRpc(key []byte, value *string) error {
	val, err := r.server.GetValue(key)
	if err != nil {
		return err
	}
	*value = val
	return nil
}
