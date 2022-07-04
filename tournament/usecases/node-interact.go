package usecases

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
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
	remote *chord.RemoteNode
	hash   func() hash.Hash
	m      uint

	// How many extra (key, value) pairs will be inserted.
	replicas uint

	ring chord.RingApi
	log  domain.Logger
}

// NewDht creates a DHT. A port is automatically designated for serving
// API. DHTs must be created in the same order in all nodes
// so communication works properly.
// Deprecated: use DhtBuilder[T] instead.
func NewDht[T any](ring chord.RingApi, data chord.DataInteract, log domain.Logger) *Dht[T] {
	return NewDhtBuilder[T]().
		Ring(ring).
		Log(log).
		Build()
}

// Get gets a value from the DHT. In case the value to retrieve is a struct,
// its serializable fields must be public.
func (dht *Dht[T]) Get(key string) (T, error) {
	var ans T
	var lastErr error = nil

	for k := range dht.withReplicas(key) {
		bkey := utils.Sha1Sized(k, dht.hash, dht.m)
		hexKey := hex.EncodeToString(bkey)

		owner, err := dht.ring.FindSuccessor(dht.remote, bkey)
		if err != nil {
			dht.log.Error(
				"error when finding key owner",
				domain.LogArgs{
					"error":   err,
					"str-key": k,
					"hex-key": hexKey,
				})
			lastErr = err
			continue
		}
		dht.log.Info(
			"Getting value...",
			domain.LogArgs{
				"str-key": k,
				"hex-key": hexKey,
				"owner":   owner.Addr(),
			})
		val, err := owner.GetValue(bkey, dht.ring)
		if err != nil {
			dht.log.Error(
				"error when getting key value",
				domain.LogArgs{
					"error":   err,
					"str-key": k,
					"hex-key": hexKey,
				})
			lastErr = err
			continue
		}
		return ans, json.Unmarshal([]byte(val), &ans)
	}
	return ans, lastErr
}

// withReplicas returns mainKey with all its replicas.
func (dht *Dht[T]) withReplicas(mainKey string) <-chan string {
	ch := make(chan string)
	go func() {
		ch <- mainKey
		for i := uint(1); i <= dht.replicas; i++ {
			ch <- fmt.Sprintf("%s_%d", mainKey, i)
		}
		close(ch)
	}()
	return ch
}

// Set sets a value in the DHT. In case the given value is a struct,
// its serializable fields must be public.
func (dht *Dht[T]) Set(key string, value T) error {
	for k := range dht.withReplicas(key) {
		bkey := utils.Sha1Sized(k, dht.hash, dht.m)

		owner, err := dht.ring.FindSuccessor(dht.remote, bkey)
		if err != nil {
			return err
		}

		bvalue, err := json.Marshal(value)
		if err != nil {
			return err
		}

		dht.log.Info(
			"Setting value...",
			domain.LogArgs{
				"str-key": k,
				"hex-key": hex.EncodeToString(bkey),
				"value":   string(bvalue),
				"owner":   owner.Addr(),
			})
		if err := owner.SetValue(bkey, string(bvalue), dht.ring); err != nil {
			return err
		}
	}
	return nil
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
	ans := []*chord.RemoteNode{dht.remote}
	node, err := dht.ring.GetSuccessor(dht.remote)
	if err != nil {
		return nil, err
	}
	for bytes.Compare(node.Id, dht.remote.Id) != 0 {
		ans = append(ans, node)

		succ, err := dht.ring.GetSuccessor(node)
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

type DhtBuilder[T any] struct {
	ring     chord.RingApi
	log      domain.Logger
	replicas uint
	remote   *chord.RemoteNode
	hashGen  func() hash.Hash
	m        uint
}

// NewDhtBuilder returns a Dht struct builder when value type is T.
// Defaults:
// m = 56,
// hashGen = sha1.New
func NewDhtBuilder[T any]() *DhtBuilder[T] {
	return &DhtBuilder[T]{
		m:       56,
		hashGen: sha1.New,
	}
}

func (builder *DhtBuilder[T]) Ring(ring chord.RingApi) *DhtBuilder[T] {
	builder.ring = ring
	return builder
}

func (builder *DhtBuilder[T]) Log(log domain.Logger) *DhtBuilder[T] {
	builder.log = log
	return builder
}

// Replicas returns a builder with the replicas field set. This is how many extra
// (key, value) pairs are stored. Default is 0 so only one (key, value) pair is stored.
func (builder *DhtBuilder[T]) Replicas(replicas uint) *DhtBuilder[T] {
	builder.replicas = replicas
	return builder
}

// M returns a builder with the m field set. This is the size of the hash in bits.
func (builder *DhtBuilder[T]) M(m uint) *DhtBuilder[T] {
	builder.m = m
	return builder
}

func (builder *DhtBuilder[T]) HashGen(hashGen func() hash.Hash) *DhtBuilder[T] {
	builder.hashGen = hashGen
	return builder
}

func (builder *DhtBuilder[T]) Remote(remote *chord.RemoteNode) *DhtBuilder[T] {
	builder.remote = remote
	return builder
}

func (builder *DhtBuilder[T]) Build() *Dht[T] {
	return &Dht[T]{
		remote:   builder.remote,
		hash:     builder.hashGen,
		m:        builder.m,
		ring:     builder.ring,
		log:      builder.log,
		replicas: builder.replicas,
	}
}
