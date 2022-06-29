package tests

import (
	"bytes"
	"crypto/sha1"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func localConfig(port uint) *chord.Config {
	return &chord.Config{
		Ip:   "127.0.0.1",
		Port: port,
		Hash: sha1.New,
		Ring: infrastruct.NewRingApi(),
		Data: infrastruct.NewDataInteract(),
	}
}

func manualId(id string, config *chord.Config) *chord.Config {
	config.Id = id
	return config
}

func TestJoin(t *testing.T) {
	entry, _ := usecases.NewNode(localConfig(8004), nil, infrastruct.NewLogger())
	node, err := usecases.NewNode(
		localConfig(8005),
		entry.RemoteNode,
		infrastruct.NewLogger(),
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 4)
	succ := node.GetSuccessor()
	if bytes.Compare(succ.Id, entry.Id) != 0 {
		t.Errorf("Expected successor of new node to be %v, got %v", entry.Addr(), succ.Addr())
	}
	pred := node.GetPredecessor()
	if bytes.Compare(pred.Id, entry.Id) != 0 {
		t.Errorf("Expected predecessor of new node to be %v, got %v", entry.Addr(), pred.Addr())
	}
	entrySucc := entry.GetSuccessor()
	if bytes.Compare(entrySucc.Id, node.Id) != 0 {
		t.Errorf("Expected successor of entry node to be %v, got %v", node.Addr(), entrySucc.Addr())
	}
	entryPred := entry.GetPredecessor()
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
	entry, _ := usecases.NewNode(localConfig(8004), nil, infrastruct.NewLogger())
	node, err := usecases.NewNode(
		localConfig(8005),
		entry.RemoteNode,
		infrastruct.NewLogger(),
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 4)

	if err := node.Ring.CheckNode(entry.RemoteNode); err != nil {
		t.Fatal(err)
	}

	if err := node.Stop(); err != nil {
		t.Fatal(err)
	}
	if err := entry.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestJoinRingOf2(t *testing.T) {
	ring1, _ := usecases.NewNode(
		manualId("1", localConfig(8004)),
		nil,
		infrastruct.NewLogger(),
	)

	ring2, err := usecases.NewNode(
		manualId("11", localConfig(8005)),
		ring1.RemoteNode,
		infrastruct.NewLogger(),
	)
	require.Nil(t, err)

	node, err := usecases.NewNode(
		manualId("21", localConfig(8006)),
		ring1.RemoteNode,
		infrastruct.NewLogger(),
	)
	require.Nil(t, err)

	time.Sleep(time.Second * 10)

	assert.Equal(t, ring1.Id, node.GetSuccessor().Id)
	assert.Equal(t, ring2.Id, node.GetPredecessor().Id)
	assert.Equal(t, node.Id, ring2.GetSuccessor().Id)
	assert.Equal(t, ring1.Id, ring2.GetPredecessor().Id)
	assert.Equal(t, ring2.Id, ring1.GetSuccessor().Id)
	assert.Equal(t, node.Id, ring1.GetPredecessor().Id)

	require.Nil(t, ring1.Stop())
	require.Nil(t, ring2.Stop())
	require.Nil(t, node.Stop())
}

func TestNodeOutFromRingOf4(t *testing.T) {
	log := infrastruct.NewLogger().ToFile()
	ring1, err := usecases.NewNode(
		manualId("1", localConfig(8004)),
		nil,
		log,
	)
	require.Nil(t, err)

	ring2, err := usecases.NewNode(
		manualId("11", localConfig(8005)),
		ring1.RemoteNode,
		log,
	)
	require.Nil(t, err)

	ring3, err := usecases.NewNode(
		manualId("21", localConfig(8006)),
		ring1.RemoteNode,
		log,
	)
	require.Nil(t, err)

	node, err := usecases.NewNode(
		manualId("31", localConfig(8007)),
		ring1.RemoteNode,
		log,
	)
	require.Nil(t, err)

	time.Sleep(time.Second * 10)

	require.Nil(t, ring2.Stop())
	time.Sleep(time.Second * 7)

	assert.Equal(t, ring1.Id, node.GetSuccessor().Id)
	assert.Equal(t, ring3.Id, node.GetPredecessor().Id)
	assert.Equal(t, ring3.Id, ring1.GetSuccessor().Id)
	assert.Equal(t, node.Id, ring1.GetPredecessor().Id)
	assert.Equal(t, node.Id, ring3.GetSuccessor().Id)
	assert.Equal(t, ring1.Id, ring3.GetPredecessor().Id)

	require.Nil(t, ring1.Stop())
	require.Nil(t, ring3.Stop())
	require.Nil(t, node.Stop())
}

func TestOneOutThenOther(t *testing.T) {
	log := infrastruct.NewLogger().ToFile()
	ring1, err := usecases.NewNode(
		manualId("1", localConfig(8004)),
		nil,
		log,
	)
	require.Nil(t, err)

	ring2, err := usecases.NewNode(
		manualId("11", localConfig(8005)),
		ring1.RemoteNode,
		log,
	)
	require.Nil(t, err)

	ring3, err := usecases.NewNode(
		manualId("21", localConfig(8006)),
		ring1.RemoteNode,
		log,
	)
	require.Nil(t, err)

	ring4, err := usecases.NewNode(
		manualId("31", localConfig(8007)),
		ring1.RemoteNode,
		log,
	)
	require.Nil(t, err)

	time.Sleep(time.Second * 10)

	require.Nil(t, ring2.Stop())
	time.Sleep(time.Second * 7) // wait for ring fixing
	require.Nil(t, ring4.Stop())

	assert.Equal(t, ring1.Id, ring3.GetSuccessor().Id)
	assert.Equal(t, ring1.Id, ring3.GetPredecessor().Id)
	assert.Equal(t, ring3.Id, ring1.GetSuccessor().Id)
	assert.Equal(t, ring3.Id, ring1.GetPredecessor().Id)

	require.Nil(t, ring1.Stop())
	require.Nil(t, ring3.Stop())
}

func TestValueSetAndGet(t *testing.T) {
	log := infrastruct.NewLogger().ToFile()

	dht := usecases.NewDht[string](
		infrastruct.NewRingApi(),
		infrastruct.NewDataInteract(),
		log)

	entry := &chord.RemoteNode{Ip: "127.0.0.1", Port: 8001}

	ring1, err := usecases.NewNode(
		manualId("1", localConfig(8002)),
		entry,
		log,
	)
	require.Nil(t, err)

	ring2, err := usecases.NewNode(
		manualId("11", localConfig(8003)),
		entry,
		log,
	)
	require.Nil(t, err)

	ring3, err := usecases.NewNode(
		manualId("21", localConfig(8004)),
		entry,
		log,
	)
	require.Nil(t, err)

	time.Sleep(time.Second * 10)

	// storing and checking a standard value
	key, value := "hello", "world"
	require.Nil(t, dht.Set(key, value))

	time.Sleep(time.Second * 3)

	val, err := dht.Get(key)
	require.Nil(t, err)
	assert.Equal(t, "world", val)

	require.Nil(t, dht.Stop())
	require.Nil(t, ring1.Stop())
	require.Nil(t, ring2.Stop())
	require.Nil(t, ring3.Stop())
}

func TestStructValueSetAndGet(t *testing.T) {
	log := infrastruct.NewLogger().ToFile()

	type test struct {
		F1 string
		F2 int
		F3 bool
	}

	dht := usecases.NewDht[test](
		infrastruct.NewRingApi(),
		infrastruct.NewDataInteract(),
		log)

	entry := &chord.RemoteNode{Ip: "127.0.0.1", Port: 8001}

	ring1, err := usecases.NewNode(
		manualId("1", localConfig(8002)),
		entry,
		log,
	)
	require.Nil(t, err)

	ring2, err := usecases.NewNode(
		manualId("11", localConfig(8003)),
		entry,
		log,
	)
	require.Nil(t, err)

	ring3, err := usecases.NewNode(
		manualId("21", localConfig(8004)),
		entry,
		log,
	)
	require.Nil(t, err)

	time.Sleep(time.Second * 10)
	k2, v2 := "struct", test{F1: "hello", F2: 1, F3: true}
	require.Nil(t, dht.Set(k2, v2))

	time.Sleep(time.Second * 3)

	val, err := dht.Get(k2)
	require.Nil(t, err)
	assert.Equal(t, v2, val)

	require.Nil(t, dht.Stop())
	require.Nil(t, ring1.Stop())
	require.Nil(t, ring2.Stop())
	require.Nil(t, ring3.Stop())
}
