package dht

import (
	"bytes"
	"crypto/sha1"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"testing"
	"time"
)

func TestJoin(t *testing.T) {
	entry, _ := NewNode(&Config{Ip: "127.0.0.1", Port: 8080, Hash: sha1.New}, nil, infrastruct.NewLogger())
	node, err := NewNode(
		&Config{Ip: "127.0.0.1", Port: 8081, Hash: sha1.New},
		entry.RemoteNode,
		infrastruct.NewLogger(),
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 4)
	succ, _ := node.GetSuccessor()
	if bytes.Compare(succ.Id, entry.Id) != 0 {
		t.Errorf("Expected successor of new node to be %v, got %v", entry.Addr(), succ.Addr())
	}
	pred, _ := node.GetPredecessor()
	if bytes.Compare(pred.Id, entry.Id) != 0 {
		t.Errorf("Expected predecessor of new node to be %v, got %v", entry.Addr(), pred.Addr())
	}
	entrySucc, _ := entry.GetSuccessor()
	if bytes.Compare(entrySucc.Id, node.Id) != 0 {
		t.Errorf("Expected successor of entry node to be %v, got %v", node.Addr(), entrySucc.Addr())
	}
	entryPred, _ := entry.GetPredecessor()
	if bytes.Compare(entryPred.Id, node.Id) != 0 {
		t.Errorf("Expected predecessor of entry node to be %v, got %v", node.Addr(), entryPred.Addr())
	}
	if err := node.Stop(); err != nil {
		t.Fatal(err)
	}
	if err := entry.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestCheckNode(t *testing.T) {
	entry, _ := NewNode(&Config{Ip: "127.0.0.1", Port: 8080, Hash: sha1.New}, nil, infrastruct.NewLogger())
	node, err := NewNode(
		&Config{Ip: "127.0.0.1", Port: 8081, Hash: sha1.New},
		entry.RemoteNode,
		infrastruct.NewLogger(),
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 4)

	if err := node.ring.CheckNode(entry.RemoteNode); err != nil {
		t.Fatal(err)
	}

	if err := node.Stop(); err != nil {
		t.Fatal(err)
	}
	if err := entry.Stop(); err != nil {
		t.Fatal(err)
	}
}
