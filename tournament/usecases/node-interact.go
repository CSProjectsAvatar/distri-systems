package usecases

import (
	"encoding/hex"
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/utils"
	"hash"
	"time"
)

// NewRemoteNode Creates an entry point to the Chord ring.
func NewRemoteNode(id string, ip string, port uint, h hash.Hash) (*chord.RemoteNode, error) {
	if _, err := h.Write([]byte(id)); err != nil {
		return nil, err
	}
	return &chord.RemoteNode{
		Id:   h.Sum(nil),
		Ip:   ip,
		Port: port,
	}, nil
}

// NewNode sets up a new node in the ring. Arg entry is the entry point to the ring.
func NewNode(config *chord.Config, entry *chord.RemoteNode, log domain.Logger) (*chord.Node, error) {
	node := &chord.Node{
		Log:  log,
		Kill: make(chan any, 1), // buffer is 1 so StopNode() can send a flag without blocking
	}

	var strId string
	if config.Id != "" {
		strId = config.Id
	} else {
		strId = fmt.Sprintf("%s:%d_%d", config.Ip, config.Port, time.Now().Unix())
	}
	var err error
	if node.RemoteNode, err = NewRemoteNode(strId, config.Ip, config.Port, config.Hash()); err != nil {
		return nil, err
	}
	log.Info(
		"Creating new node...",
		domain.LogArgs{
			"id": hex.EncodeToString(node.Id),
		})

	m := config.Hash().Size() * 8

	node.Ftable = newFtable(node.RemoteNode, m)

	node.Ring = config.Ring
	node.Data = config.Data

	node.Ring.StartNode(node) // now node is a ring server

	node.PredMtx.Lock()
	node.Predecessor = nil
	node.PredMtx.Unlock()

	if entry != nil {
		if err := node.Join(entry); err != nil {
			return nil, err
		}
	} else {
		// create() implementation from paper
		node.SuccMtx.Lock()
		node.Successor = node.RemoteNode
		node.SuccMtx.Unlock()
	}
	go utils.RepeatAction(node.Stabilize, 1*time.Second, node.Kill)         // stabilize node once in a second
	go ensureFtable(node, m)                                                // ensure finger table is ok
	go utils.RepeatAction(node.CheckPredecessor, 10*time.Second, node.Kill) // check predecessor is alive once in 10 seconds

	return node, nil
}

func ensureFtable(node *chord.Node, m int) {
	next := 0
	t := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-t.C:
			next = node.UpdateFingerRow(next, m)
		case <-node.Kill:
			t.Stop()
			return
		}
	}
}
