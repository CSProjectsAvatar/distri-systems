package dht

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"hash"
	"sync"
)

type RemoteNode struct {
	Id   []byte
	Ip   string
	Port uint
}

type Node struct {
	*RemoteNode
	ring        RingApi
	ftable      ftable
	successor   *RemoteNode
	kill        chan any // kill is used to know when to stop the node
	log         usecases.Logger
	predecessor *RemoteNode
	succMtx     sync.RWMutex
	predMtx     sync.RWMutex
	ftableMtx   sync.RWMutex
	data        DataInteract
}

type Config struct {
	Id   string
	Ip   string
	Port uint
	Hash func() hash.Hash
}

// Finger Table.
type ftable []*fingerRow

// A row of the Finger Table.
type fingerRow struct {
	Id   []byte      // hash of (n + 2^i) mod 2^m (i>=0)
	Node *RemoteNode // first node on circle that succeeds Id
}

// Data is a key-value pair of string.
type Data struct {
	Key, Value string
}
