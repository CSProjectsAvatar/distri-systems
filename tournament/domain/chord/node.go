package chord

import (
	"bytes"
	"sync"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/utils"
)

type Node struct {
	*RemoteNode
	Ring        RingApi
	Ftable      Ftable
	Successor   *RemoteNode
	Kill        chan any // Kill is used to know when to stop the node
	Log         domain.Logger
	Predecessor *RemoteNode
	SuccMtx     sync.RWMutex
	PredMtx     sync.RWMutex
	FtableMtx   sync.RWMutex
	Data        DataInteract
}

// CheckPredecessor checks for the existence of the predecessor of node. If it doesn't exist, then a nil value is set.
func (node *Node) CheckPredecessor() {
	node.PredMtx.RLock()
	pred := node.Predecessor
	node.PredMtx.RUnlock()

	if pred != nil {
		if err := node.Ring.CheckNode(pred); err != nil {
			node.Log.Error(
				"Predecessor is out.",
				domain.LogArgs{
					"me":   node.Addr(),
					"pred": pred.Addr(),
				},
			)
			node.PredMtx.Lock()
			node.Predecessor = nil
			node.PredMtx.Unlock()
		}
	}

}

// Stabilize updates successor of node.
func (node *Node) Stabilize() {
	node.SuccMtx.RLock()
	succ := node.Successor
	if succ == nil {
		node.SuccMtx.RUnlock()
		panic("successor is nil")
	}
	node.SuccMtx.RUnlock()

	succPred, err := succ.getPredecessor(node.Ring)
	if err != nil {
		node.Log.Error(
			"Predecessor retrieval failed when stabilizing.",
			domain.LogArgs{
				"error":     err,
				"me":        node.Addr(),
				"successor": succ.Addr(),
			},
		)
		// todo @audit it's better to change successor to a finger rather than to itself
		succ = node.RemoteNode // now I am my own successor so my successor will be properly updated in next calls

		node.PredMtx.RLock()
		succPred = node.Predecessor
		node.PredMtx.RUnlock()
	}
	if unvoid(succPred) && utils.InInterval(succPred.Id, node.Id, succ.Id) {
		node.SuccMtx.Lock()
		node.Log.Info(
			"Updating successor.",
			domain.LogArgs{
				"host":           node.RemoteNode.ID(),
				"prev successor": node.Successor.ID(),
				"new successor":  succPred.ID()})
		node.Successor = succPred
		succ = succPred
		node.SuccMtx.Unlock()
	}
	if bytes.Compare(node.Id, succ.Id) != 0 { // node != succ
		node.Log.Trace(
			"notifying...",
			domain.LogArgs{
				"me":        node.Addr(),
				"successor": succ.Addr(),
			})
		if err := succ.notify(node.RemoteNode, node.Ring); err != nil {
			node.Log.Error(
				"Error when notifying.",
				domain.LogArgs{
					"node":      node.Addr(),
					"successor": succ.Addr(),
					"error":     err,
				},
			)
		}
	}
}

// Join puts node inside ring as a predecessor of entry.
func (node *Node) Join(entry *RemoteNode) error {
	if entry == nil {
		panic("entry point is nil. At least one node of the ring must be known.")
	}
	succEntry, err := entry.findSuccessor(node.Id, node.Ring)
	if err != nil {
		return err
	}
	if bytes.Compare(node.Id, succEntry.Id) == 0 {
		return ErrNodeAlreadyExists
	}
	node.SuccMtx.Lock()

	node.Log.Info(
		"Updating successor.",
		domain.LogArgs{
			"host":           node.RemoteNode.ID(),
			"prev successor": node.Successor.ID(),
			"new successor":  succEntry.ID()})

	node.Successor = succEntry
	node.SuccMtx.Unlock()

	return nil
}

// fix_fingers() method from paper.
func (node *Node) UpdateFingerRow(row int, m uint) int {
	id := FingerId(node.Id, row, m)
	succEntry, err := node.findSuccessor(id)
	nextRow := (row + 1) % int(m)

	if err != nil || succEntry == nil {
		node.Log.Error(
			"error when finding finger.",
			domain.LogArgs{
				"error":    err,
				"row hash": id,
				"me":       node.Addr(),
			},
		)
		// @todo handle retry, passing ahead for now

		return nextRow
	}
	node.FtableMtx.Lock()

	newf := NewFingerRow(id, succEntry)
	if node.Ftable[row] == nil || bytes.Compare(newf.Node.Id, node.Ftable[row].Node.Id) != 0 {
		var prev string
		if node.Ftable[row] != nil {
			prev = node.Ftable[row].Node.Addr()
		}
		node.Log.Trace(
			"changing finger...",
			domain.LogArgs{
				"prev finger": prev,
				"new finger":  newf.Node.Addr(),
				"row":         row,
				"me":          node.Addr(),
			},
		)
	}
	node.Ftable[row] = newf

	node.FtableMtx.Unlock()

	return nextRow
}

func (node *Node) findSuccessor(id []byte) (*RemoteNode, error) {
	curr := node.RemoteNode

	node.SuccMtx.RLock()
	defer node.SuccMtx.RUnlock()
	succ := node.Successor

	if succ == nil {
		panic("successor is nil")
	}
	if utils.InIntervalRIncluded(id, curr.Id, succ.Id) { // id ∈ (n, successor]
		return succ, nil
	}
	toAsk := node.closestPrecedingNode(id)
	node.Log.Trace(
		"closest preceding node",
		domain.LogArgs{
			"addr": toAsk.Addr(),
			"me":   node.Addr(),
		},
	)

	if bytes.Compare(toAsk.Id, node.Id) == 0 {
		succ := node.GetSuccessor()
		if succ == nil {
			// return toAsk, nil
			panic("successor is nil")
		}
		return succ, nil
	}
	var err error
	succ, err = toAsk.findSuccessor(id, node.Ring)
	if err != nil {
		return nil, err
	}
	if succ == nil {
		panic("successor is nil")
	}
	return succ, nil
}

func (node *Node) closestPrecedingNode(id []byte) *RemoteNode {
	node.FtableMtx.RLock()
	defer node.FtableMtx.RUnlock()

	n := node.RemoteNode
	for i := len(node.Ftable) - 1; i >= 0; i-- {
		frow := node.Ftable[i]
		if frow == nil {
			continue
		}
		if utils.InInterval(frow.Node.Id, n.Id, id) {
			return frow.Node
		}
	}
	return n
}

func (node *Node) Stop() error {
	close(node.Kill)
	if err := node.Ring.StopNode(); err != nil {
		return err
	}
	return node.Data.Close()
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

func (node *Node) GetSuccessor() *RemoteNode {
	node.SuccMtx.RLock()
	defer node.SuccMtx.RUnlock()
	return node.Successor
}

// Notify updates the node predecessor.
func (node *Node) Notify(pred *RemoteNode) error {
	node.PredMtx.Lock()
	defer node.PredMtx.Unlock()

	if node.Predecessor == nil || utils.InInterval(pred.Id, node.Predecessor.Id, node.Id) {
		node.Log.Info(
			"Updating predecessor.",

			domain.LogArgs{
				"node":             node.ID(),
				"prev predecessor": node.Predecessor.ID(),
				"new predecessor":  pred.ID()})
		node.Predecessor = pred

		node.SuccMtx.RLock()
		if bytes.Compare(node.Predecessor.Id, node.Successor.Id) != 0 {
			if err := node.moveKeysPred(node.Predecessor); err != nil {
				return err
			}
		}
		node.SuccMtx.RUnlock()
	}
	return nil
}

// moveKeysPred moves all keys less than newPred.ID in node to newPred. This is called when new predecessor arrives.
func (node *Node) moveKeysPred(newPred *RemoteNode) error {
	predData, err := node.Data.LowerEq(newPred.Id)
	if err != nil {
		return err
	}
	if predData != nil && len(predData) > 0 {
		node.Log.Info(
			"moving data",
			domain.LogArgs{
				"from": node.Addr(),
				"to":   newPred.Addr(),
				//"data":        hex.EncodeToString(predData[0].Key),
				"keys-amount": len(predData),
			},
		)
	}
	if err := node.Ring.SendData(predData, newPred); err != nil {
		return err
	}
	return node.Data.Delete(predData)
}

func (node *Node) GetPredecessor() *RemoteNode {
	node.PredMtx.RLock()
	defer node.PredMtx.RUnlock()
	return node.Predecessor
}

func (node *Node) ReceiveData(data []*Data) error {
	return node.Data.Save(data)
}

func (node *Node) GetValue(key []byte) (string, error) {
	return node.Data.Get(key)
}

func (node *Node) SetValue(key []byte, value string) error {
	return node.Data.Save([]*Data{{key, value}})
}
