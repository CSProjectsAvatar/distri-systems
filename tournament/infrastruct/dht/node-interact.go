package dht

import (
	"bytes"
	"hash"
	"math/big"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/CSProjectsAvatar/distri-systems/utils"
)

// NewRemoteNode Creates an entry point to the Chord ring.
func NewRemoteNode(id string, addr string, h hash.Hash) (*RemoteNode, error) {
	if _, err := h.Write([]byte(id)); err != nil {
		return nil, err
	}
	return &RemoteNode{
		Id:   h.Sum(nil),
		Addr: addr,
	}, nil
}

func NewNode(config *Config, entry *RemoteNode, log usecases.Logger) (*Node, error) {
	node := &Node{
		log:  log,
		kill: make(chan any, 1), // buffer is 1 so StopNode() can send a flag without blocking
	}

	var strId string
	if config.Id != "" {
		strId = config.Id
	} else {
		strId = config.Addr
	}
	var err error
	if node.RemoteNode, err = NewRemoteNode(strId, config.Addr, config.Hash()); err != nil {
		return nil, err
	}
	log.Info("Creating new node with ID = %d\n", new(big.Int).SetBytes(node.RemoteNode.Id))

	m := config.Hash().Size() * 8

	node.ftable = newFtable(node.RemoteNode, m)

	ring, err := newRingApi(config)
	if err != nil {
		return nil, err
	}
	node.ring = ring
	node.ring.StartNode(node) // now node is a ring server

	if err := node.join(entry); err != nil {
		return nil, err
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
			node.log.Error(
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
	if err != nil || succPred == nil {
		node.log.Error("predecessor retrieval failed when stabilizing.\n"+
			"\tError: %v "+
			"\tPredecessor of successor: %v",
			err, succPred,
		)
		return
	}
	if succPred.Id != nil && utils.InInterval(succPred.Id, node.Id, succ.Id) { // @audit why can succPred.Id be nil?
		node.succMtx.Lock()
		node.successor = succPred
		node.succMtx.Unlock()
	}
	if err := succ.notify(node.RemoteNode, node.ring); err != nil {
		node.log.Error(
			"error when node with ID = %v was notifying successor with ID = %v.",
			node.Id, succ.Id,
		)
	}
}

// join puts node inside ring as a predecessor of entry.
func (node *Node) join(entry *RemoteNode) error {
	nodeToAsk := entry
	if entry == nil {
		nodeToAsk = node.RemoteNode
	}
	succEntry, err := nodeToAsk.findSuccessor(node.Id, node.ring)
	if err != nil {
		return err
	}
	if bytes.Compare(node.Id, succEntry.Id) == 0 {
		return ErrNodeAlreadyExists
	}
	node.succMtx.Lock()
	node.successor = succEntry
	node.succMtx.Unlock()

	return nil
}

func newRingApi(config *Config) (RingApi, error) {
	panic("implement me")
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
		node.log.Error(
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
		succ, err = toAsk.getSuccessor(node.ring)
		if err != nil {
			return nil, err
		}
		if succ == nil {
			return toAsk, nil
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
	succ, err := node.findSuccessor(id)
	if err != nil {
		return nil, err
	}
	if succ == nil {
		return nil, ErrNodeNotFound
	}
	return succ, nil
}

func (node *Node) GetSuccessor() (*RemoteNode, error) {
	node.succMtx.RLock()
	defer node.succMtx.RUnlock()
	return node.successor, nil
}

// Notify updates the node predecessor.
func (node *Node) Notify(pred *RemoteNode) error {
	node.predMtx.Lock()
	defer node.predMtx.Unlock()

	if node.predecessor == nil || utils.InInterval(pred.Id, node.predecessor.Id, node.Id) {
		node.log.Info("updating predecessor of %v to %v", node.Id, pred.Id)
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

// @todo ReceiveData() (endpoint of SendData RPC method)
func (node *Node) ReceiveData(data []*Data) error {

	panic("not implemented")
}
