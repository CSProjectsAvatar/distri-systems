package chord

import (
	"encoding/hex"
	"fmt"
)

type RemoteNode struct {
	Id   []byte
	Ip   string
	Port uint
}

// Helper for RingApi.Notify().
func (node *RemoteNode) notify(pred *RemoteNode, ring RingApi) error {
	return ring.Notify(node, pred)
}

// Helper for RingApi.GetPredecessor().
func (node *RemoteNode) getPredecessor(ring RingApi) (*RemoteNode, error) {
	return ring.GetPredecessor(node)

}

// Helper for RingApi.FindSuccessor().
func (node *RemoteNode) findSuccessor(id []byte, ring RingApi) (*RemoteNode, error) {
	return ring.FindSuccessor(node, id)
}

// Helper for RingApi.GetSuccessor().
func (node *RemoteNode) getSuccessor(ring RingApi) (*RemoteNode, error) {
	return ring.GetSuccessor(node)
}

func (node *RemoteNode) Addr() string {
	if node == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s:%d", node.Ip, node.Port)
}
func (node *RemoteNode) ID() string {
	if node == nil {
		return "<nil>"
	}
	return hex.EncodeToString(node.Id)
}

func (node *RemoteNode) SetValue(key []byte, value string, ring RingApi) error {
	return ring.SetValue(node, key, value)
}

func (node *RemoteNode) GetValue(key []byte, ring RingApi) (string, error) {
	return ring.GetValue(node, key)
}

func unvoid(node *RemoteNode) bool {
	return node != nil && node.Id != nil && len(node.Id) != 0
}
