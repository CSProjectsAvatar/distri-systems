package usecases

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/utils"
	"hash"
	"time"
)

// NewRemoteNode creates an entry point to the Chord ring. The given hash function is used
// to hash the node ID. If h is nil, then no hash is applied.
func NewRemoteNode(id string, ip string, port uint, h func() hash.Hash, m uint) (*chord.RemoteNode, error) {
	ans := &chord.RemoteNode{
		Ip:   ip,
		Port: port,
	}
	if h == nil {
		ans.Id = make([]byte, m/8)
		copy(ans.Id, id)
	} else {
		ans.Id = utils.Sha1Sized(id, h, m)
	}
	return ans, nil
}

// NewNode sets up a new node in the ring. Arg entry is the entry point to the ring.
func NewNode(config *chord.Config, entry *chord.RemoteNode, log domain.Logger) (*chord.Node, error) {
	node := &chord.Node{
		Log:  log,
		Kill: make(chan any, 1), // buffer is 1 so StopNode() can send a flag without blocking
	}

	var strId string
	var hashFactory func() hash.Hash = nil
	if config.Id != nil {
		strId = string(config.Id)
	} else {
		strId = fmt.Sprintf("%s:%d", config.Ip, config.Port)
		if config.IncludeDate {
			strId = fmt.Sprintf("%s_%d", strId, time.Now().Unix())
		}
		hashFactory = config.Hash
	}
	var err error
	if node.RemoteNode, err = NewRemoteNode(strId, config.Ip, config.Port, hashFactory, config.M); err != nil {
		return nil, err
	}
	log.Info(
		"Creating new node...",
		domain.LogArgs{
			"id":   hex.EncodeToString(node.Id),
			"addr": node.Addr(),
		})

	node.Ftable = newFtable(node.RemoteNode, config.M)

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
	go ensureFtable(node, config.M)                                         // ensure finger table is ok
	go utils.RepeatAction(node.CheckPredecessor, 10*time.Second, node.Kill) // check predecessor is alive once in 10 seconds

	return node, nil
}

func ensureFtable(node *chord.Node, m uint) {
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

type Dht[T any] struct {
	chord *chord.Node
	hash  func() hash.Hash
	m     uint
}

var port uint = 8001

func NewTestDht[T any]() *Dht[T] {
	return newDht[T](
		infrastruct.NewRingApi(),
		infrastruct.NewNamedDataInteract(fmt.Sprintf("bunt-8001-%v", time.Now())),
		infrastruct.NewLogger().ToFile(),
		false)
}

// NewDht creates a DHT. A port is automatically designated for serving
// API. DHTs must be created in the same order in all nodes
// so communication works properly.
func NewDht[T any](ring chord.RingApi, data chord.DataInteract, log domain.Logger) *Dht[T] {
	return newDht[T](ring, data, log, true)
}

func newDht[T any](ring chord.RingApi, data chord.DataInteract, log domain.Logger, includeDate bool) *Dht[T] {
	var entry *chord.RemoteNode = nil // result of discovering policy
	hashGen := sha1.New
	m := uint(56)

	node, err := NewNode(
		&chord.Config{
			Ip:          "127.0.0.1",
			Port:        port,
			Hash:        hashGen,
			Ring:        ring,
			Data:        data,
			M:           m,
			IncludeDate: includeDate,
		},
		entry,
		log)

	if err != nil {
		panic(err)
	}
	port++

	return &Dht[T]{
		chord: node,
		hash:  hashGen,
		m:     m,
	}
}

// Get gets a value from the DHT. In case the value to retrieve is a struct,
// its serializable fields must be public.
func (dht *Dht[T]) Get(key string) (T, error) {
	var ans T

	bkey := utils.Sha1Sized(key, dht.hash, dht.m)

	owner, err := dht.chord.FindSuccessor(bkey)
	if err != nil {
		return ans, err
	}

	dht.chord.Log.Info(
		"Getting value...",
		domain.LogArgs{
			"str-key": key,
			"hex-key": hex.EncodeToString(bkey),
			"owner":   owner.Addr(),
		})
	val, err := owner.GetValue(bkey, dht.chord.Ring)
	if err != nil {
		return ans, err
	}
	return ans, json.Unmarshal([]byte(val), &ans)
}

// Set sets a value in the DHT. In case the given value is a struct,
// its serializable fields must be public.
func (dht *Dht[T]) Set(key string, value T) error {
	bkey := utils.Sha1Sized(key, dht.hash, dht.m)

	owner, err := dht.chord.FindSuccessor(bkey)
	if err != nil {
		return err
	}

	bvalue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	dht.chord.Log.Info(
		"Setting value...",
		domain.LogArgs{
			"str-key": key,
			"hex-key": hex.EncodeToString(bkey),
			"value":   string(bvalue),
			"owner":   owner.Addr(),
		})
	return owner.SetValue(bkey, string(bvalue), dht.chord.Ring)
}

func (dht *Dht[T]) Stop() error {
	return dht.chord.Stop()
}

// RingList returns the ring in a clockwise list of node addresses.
func (dht *Dht[T]) RingList() ([]string, error) {
	nodes, err := dht.allNodes()
	if err != nil {
		return nil, err
	}
	var nodesStr []string
	for _, n := range nodes {
		nodesStr = append(nodesStr, n.Addr())
	}
	return nodesStr, nil
}

func (dht *Dht[T]) allNodes() ([]*chord.RemoteNode, error) {
	ans := []*chord.RemoteNode{dht.chord.RemoteNode}
	for node := dht.chord.GetSuccessor(); bytes.Compare(node.Id, dht.chord.Id) != 0; {
		ans = append(ans, node)

		succ, err := dht.chord.Ring.GetSuccessor(node)
		if err != nil {
			return nil, err
		}
		if bytes.Compare(succ.Id, node.Id) == 0 || node.Id == nil || len(node.Id) == 0 {
			return nil, fmt.Errorf("ring is not closed yet")
		}
		node = succ
	}
	return ans, nil
}
