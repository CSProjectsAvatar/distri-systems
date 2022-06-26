package chord

import (
	"errors"
	"hash"
)

var ErrNodeAlreadyExists = errors.New("node already exists")
var ErrNodeNotFound = errors.New("node not found")

type Config struct {
	Id   string
	Ip   string
	Port uint
	Hash func() hash.Hash
	Ring RingApi
	Data DataInteract
}

// Data is a key-value pair of string.
type Data struct {
	Key, Value string
}
