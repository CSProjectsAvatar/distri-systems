package tests

import (
	"testing"
	// . "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	// use "github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	//mock
	mock "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/mocking"
)

func TestGiveRightLeaderToTheNewOne(t *testing.T) {
	log.Println("> Election - Not in List <")
	assert := assert.New(t)

	server1Addr := "localhost:40051"
	server2Addr := "localhost:40052"

	node1 := mock.NewMockRoutine(server1Addr, server1Addr)
	node2 := mock.NewMockRoutine(server2Addr, server1Addr)

	assert.Equal(server1Addr, node2.Elect.GetLeader())
	assert.NotNil(node1)
}
func TestElectionChangeLeader(t *testing.T) {
	log.Println("> Election - Not in List <")
	assert := assert.New(t)
	require := require.New(t)

	server1Addr := "localhost:50051"
	server2Addr := "localhost:50052"

	node1 := mock.NewMockRoutine(server1Addr, server1Addr)
	node2 := mock.NewMockRoutine(server2Addr, server1Addr)
	node1.SetSuccessor(server2Addr)

	assert.NotNil(node1)
	assert.Equal(server1Addr, node2.Elect.GetLeader())
	assert.Equal(server1Addr, node1.Elect.GetLeader())

	node1.Elect.CreateElection()

	//time.Sleep(time.Second * 30)
	<-node1.Elect.OnLeaderChange()
	require.Equal(server2Addr, node1.Elect.GetLeader())
	require.Equal(server2Addr, node2.Elect.GetLeader())

	// node 2 is off
	node1.SetSuccessor(server1Addr)

	node1.Elect.CreateElection()

	<-node1.Elect.OnLeaderChange()
	require.Equal(server1Addr, node2.GetSuccessor())
	require.Equal(server1Addr, node1.GetSuccessor())

	require.Equal(server1Addr, node1.Elect.GetLeader())
	require.Equal(server2Addr, node2.Elect.GetLeader())
}
func TestElectionChangeLeader4(t *testing.T) {
	log.Println("> Election - Not in List <")
	assert := assert.New(t)
	require := require.New(t)

	server1Addr := "localhost:50151"
	server2Addr := "localhost:50252"
	server3Addr := "localhost:50353"
	server4Addr := "localhost:50454"

	node1 := mock.NewMockRoutine(server1Addr, server1Addr)
	node4 := mock.NewMockRoutine(server4Addr, server1Addr)
	node3 := mock.NewMockRoutine(server3Addr, server4Addr)
	node2 := mock.NewMockRoutine(server2Addr, server3Addr)
	node1.SetSuccessor(server2Addr)

	assert.NotNil(node1)
	assert.Equal(server1Addr, node2.Elect.GetLeader())
	assert.Equal(server1Addr, node3.Elect.GetLeader())
	assert.Equal(server1Addr, node4.Elect.GetLeader())
	assert.Equal(server1Addr, node1.Elect.GetLeader())

	node2.Elect.CreateElection()

	<-node1.Elect.OnLeaderChange()
	<-node2.Elect.OnLeaderChange()
	<-node3.Elect.OnLeaderChange()
	<-node4.Elect.OnLeaderChange()
	assert.Equal(server4Addr, node1.Elect.GetLeader())
	assert.Equal(server4Addr, node2.Elect.GetLeader())
	assert.Equal(server4Addr, node3.Elect.GetLeader())
	assert.Equal(server4Addr, node4.Elect.GetLeader())

	// node 4 is off
	node3.SetSuccessor(server1Addr)

	node1.Elect.CreateElection()

	<-node1.Elect.OnLeaderChange()
	<-node2.Elect.OnLeaderChange()
	<-node3.Elect.OnLeaderChange()
	require.Equal(server1Addr, node3.GetSuccessor())
	require.Equal(server3Addr, node2.GetSuccessor())
	require.Equal(server2Addr, node1.GetSuccessor())

	require.Equal(server3Addr, node1.Elect.GetLeader())
	require.Equal(server3Addr, node2.Elect.GetLeader())
	require.Equal(server3Addr, node3.Elect.GetLeader())
}
