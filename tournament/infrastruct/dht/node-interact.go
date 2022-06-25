package dht

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"hash"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/CSProjectsAvatar/distri-systems/utils"
)

// NewRemoteNode Creates an entry point to the Chord ring.
func NewRemoteNode(id string, ip string, port uint, h hash.Hash) (*RemoteNode, error) {
	if _, err := h.Write([]byte(id)); err != nil {
		return nil, err
	}
	return &RemoteNode{
		Id:   h.Sum(nil),
		Ip:   ip,
		Port: port,
	}, nil
}

// NewNode sets up a new node in the ring. Arg entry is the entry point to the ring.
func NewNode(config *Config, entry *RemoteNode, log usecases.Logger) (*Node, error) {
	node := &Node{
		log:  log,
		kill: make(chan any, 1), // buffer is 1 so StopNode() can send a flag without blocking
	}

	var strId string
	if config.Id != "" {
		strId = config.Id
	} else {
		strId = config.Ip + ":" + fmt.Sprint(config.Port)
	}
	var err error
	if node.RemoteNode, err = NewRemoteNode(strId, config.Ip, config.Port, config.Hash()); err != nil {
		return nil, err
	}
	log.Info(
		"Creating new node...",
		usecases.LogArgs{
			"id": hex.EncodeToString(node.Id),
		})

	m := config.Hash().Size() * 8

	node.ftable = newFtable(node.RemoteNode, m)

	ring, err := newRingApi(config)
	if err != nil {
		return nil, err
	}
	node.ring = ring
	node.ring.StartNode(node) // now node is a ring server

	node.predMtx.Lock()
	node.predecessor = nil
	node.predMtx.Unlock()

	if entry != nil {
		if err := node.join(entry); err != nil {
			return nil, err
		}
	} else {
		// create() implementation from paper
		node.succMtx.Lock()
		node.successor = node.RemoteNode
		node.succMtx.Unlock()
	}
	go utils.RepeatAction(node.stabilize, 1*time.Second, node.kill)         // stabilize node once in a second
	go ensureFtable(node, m)                                                // ensure finger table is ok
	go utils.RepeatAction(node.checkPredecessor, 10*time.Second, node.kill) // check predecessor is alive once in 10 seconds

	return node, nil
}

// checkPredecessor checks for the existence of the predecessor of node. If it doesn't exist, then a nil value is set.
func (node *Node) checkPredecessor() {
	node.predMtx.RLock()
	pred := node.predecessor
	node.predMtx.RUnlock()

	if pred != nil {
		if err := node.ring.CheckNode(pred); err != nil {
			node.log.Errorf(
				"node with ID = %v failed. Message sent from node %v.",
				pred.Id, node.Id,
			)
			node.predMtx.Lock()
			node.predecessor = nil
			node.predMtx.Unlock()
		}
	}

}

// stabilize updates successor of node.
func (node *Node) stabilize() {
	node.succMtx.RLock()
	succ := node.successor
	if succ == nil {
		node.succMtx.RUnlock()
		return
	}
	node.succMtx.RUnlock()

	succPred, err := succ.getPredecessor(node.ring)
	if err != nil {
		node.log.Errorf("predecessor retrieval failed when stabilizing.\n"+
			"\tError: %v "+
			"\tPredecessor of successor: %v",
			err, succPred,
		)
		return
	}
	if unvoid(succPred) && utils.InInterval(succPred.Id, node.Id, succ.Id) {
		node.succMtx.Lock()
		node.log.Info(
			"Updating successor.",
			usecases.LogArgs{
				"host":           node.RemoteNode.Addr(),
				"prev successor": node.successor.Addr(),
				"new successor":  succPred.Addr()})
		node.successor = succPred
		node.succMtx.Unlock()
	}
	if bytes.Compare(node.Id, succ.Id) != 0 { // node != succ
		if err := succ.notify(node.RemoteNode, node.ring); err != nil {
			node.log.Error(
				"Error when notifying.",
				usecases.LogArgs{
					"node":      node.Addr(),
					"successor": succ.Addr(),
					"error":     err,
				},
			)
		}
	}
}

// join puts node inside ring as a predecessor of entry.
func (node *Node) join(entry *RemoteNode) error {
	if entry == nil {
		panic("entry point is nil. At least one node of the ring must be known.")
	}
	succEntry, err := entry.findSuccessor(node.Id, node.ring)
	if err != nil {
		return err
	}
	if bytes.Compare(node.Id, succEntry.Id) == 0 {
		return ErrNodeAlreadyExists
	}
	node.succMtx.Lock()
	node.log.Info(
		"Updating successor.",
		usecases.LogArgs{
			"host":           node.RemoteNode.Addr(),
			"prev successor": node.successor.Addr(),
			"new successor":  succEntry.Addr()})
	node.successor = succEntry
	node.succMtx.Unlock()

	return nil
}

func newRingApi(config *Config) (RingApi, error) {
	return &RpcRing{}, nil
}

func ensureFtable(node *Node, m int) {
	next := 0
	t := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-t.C:
			next = node.updateFingerRow(next, m)
		case <-node.kill:
			t.Stop()
			return
		}
	}
}

// fix_fingers() method from paper.
func (node *Node) updateFingerRow(row int, m int) int {
	id := fingerId(node.Id, row, m)
	succEntry, err := node.findSuccessor(id)
	nextRow := (row + 1) % m

	if err != nil || succEntry == nil {
		node.log.Errorf(
			"error when finding finger. "+
				"\n\tError: %v "+
				"\n\tRemote Node: %v"+
				"\n\tHost ID: %x"+
				"\n\trow hash: %x",
			err, succEntry, node.Id, id,
		)
		// @todo handle retry, passing ahead for now

		return nextRow
	}
	node.ftableMtx.Lock()
	node.ftable[row] = newFingerRow(id, succEntry)
	node.ftableMtx.Unlock()

	return nextRow
}

func (node *Node) findSuccessor(id []byte) (*RemoteNode, error) {
	curr := node.RemoteNode

	node.succMtx.RLock()
	defer node.succMtx.RUnlock()
	succ := node.successor

	if succ == nil {
		return curr, nil
	}
	if utils.InIntervalRIncluded(id, curr.Id, succ.Id) { // id ∈ (n, successor]
		return succ, nil
	}
	toAsk := node.closestPrecedingNode(id)
	var err error
	if bytes.Compare(toAsk.Id, node.Id) == 0 {
		succ, err := node.GetSuccessor()
		if err != nil {
			return nil, err
		}
		if succ == nil {
			// return toAsk, nil
			panic("successor is nil")
		}
		return succ, nil
	}
	succ, err = toAsk.findSuccessor(id, node.ring)
	if err != nil {
		return nil, err
	}
	if succ == nil {
		return curr, nil
	}
	return succ, nil
}

func (node *Node) closestPrecedingNode(id []byte) *RemoteNode {
	// @audit check if lock for node.predecessor is needed
	node.predMtx.RLock()
	defer node.predMtx.RUnlock()

	n := node.RemoteNode
	for i := len(node.ftable) - 1; i >= 0; i-- {
		frow := node.ftable[i]
		if frow == nil || frow.Node == nil {
			continue
		}
		if utils.InInterval(frow.Id, n.Id, id) {
			return frow.Node
		}
	}
	return n
}

func (node *Node) Stop() error {
	node.kill <- 0
	return node.ring.StopNode()
}

func (node *Node) FindSuccessor(id []byte) (*RemoteNode, error) {
	node.log.Info(
		"FindSuccessor RPC.",
		usecases.LogArgs{
			"server": node.RemoteNode.Addr(),
			"id":     hex.EncodeToString(id)})

	succ, err := node.findSuccessor(id)
	if err != nil {
		return nil, err
	}
	if succ == nil {
		return nil, ErrNodeNotFound
	}
	return succ, nil
}

func (node *Node) GetSuccessor() (*RemoteNode, error) { // todo @audit this method is returning an error in vane
	node.succMtx.RLock()
	defer node.succMtx.RUnlock()
	return node.successor, nil
}

// Notify updates the node predecessor.
func (node *Node) Notify(pred *RemoteNode) error {
	node.predMtx.Lock()
	defer node.predMtx.Unlock()

	if node.predecessor == nil || utils.InInterval(pred.Id, node.predecessor.Id, node.Id) {
		node.log.Info(
			"Updating predecessor.",
			usecases.LogArgs{
				"node":             node.Addr(),
				"prev predecessor": node.predecessor.Addr(),
				"new predecessor":  pred.Addr()})
		prevPred := node.predecessor
		node.predecessor = pred

		if prevPred != nil {
			if err := node.moveKeysPred(node.predecessor); err != nil {
				return err
			}
		}
	}
	return nil
}

// moveKeysPred moves all keys less than newPred.ID in node to newPred. This is called when new predecessor arrives.
func (node *Node) moveKeysPred(newPred *RemoteNode) error {
	predData := node.data.LowerEq(newPred.Id)
	if err := node.ring.SendData(predData, newPred); err != nil {
		return err
	}
	node.data.Delete(predData)

	return nil
}

func (node *Node) GetPredecessor() (*RemoteNode, error) {
	node.predMtx.RLock()
	defer node.predMtx.RUnlock()
	return node.predecessor, nil
}

func (node *Node) ReceiveData(data []*Data) error {
	node.data.Save(data)
	return nil
}
