package dht

import pb "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_chord"

type ChordTransport interface { // Used on Chord Algo.
	Start() error
	Stop() error

	//RPC
	GetSuccessor(*Node) (*Node, error)
	FindSuccessor(*Node, []byte) (*Node, error)
	GetPredecessor(*Node) (*Node, error)
	Notify(*Node, *Node) error
	CheckPredecessor(*Node) error
	SetPredecessor(*Node, *Node) error
	SetSuccessor(*Node, *Node) error

	//Storage
	GetKey(*Node, string) (*pb.GetResponse, error)
	SetKey(*Node, string, string) error
	DeleteKey(*Node, string) error
	RequestKeys(*Node, []byte, []byte) ([]*KV, error)
	DeleteKeys(*Node, []string) error
}
