package chord

import (
	"errors"
	"hash"
)

var (
	ErrNodeAlreadyExists = errors.New("node already exists")
	ErrNodeNotFound      = errors.New("node not found")
	ErrKeyExists         = errors.New("key exists")
	ErrKeyNotFound       = errors.New("key not found")
)

type Config struct {
	Id   string
	Ip   string
	Port uint
	Hash func() hash.Hash
	Ring RingApi
	Data DataInteract

	// Identifiers length.
	M uint
}

// Data is a key-value pair of string.
type Data struct {
	Key   []byte
	Value string
}
