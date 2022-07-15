package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	inter "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// Build Chord Node
func BuildChordNode(remote *chord.RemoteNode, logger domain.Logger) *chord.Node {
	localIp := "127.0.0.1"
	entry, err := usecases.NewNode(
		ChordConfig(localIp, domain.ChordPort),
		remote,
		logger)
	if err != nil {
		log.Fatal("Couldn't intialize chord", err) // I know, tell me later
	}
	return entry
}

// Build DataMngr
func BuildDataMngr(chordSrvIp string, chordSrvPort uint) usecases.DataMngr {
	remote := &chord.RemoteNode{
		Ip:   chordSrvIp,
		Port: chordSrvPort,
	}

	dhtStr := NewTestDht[string](domain.ReplicaNumber, remote)
	dhtInfos := NewTestDht[[]*domain.TournInfo](domain.ReplicaNumber, remote)
	dhtMatches := NewTestDht[[]*domain.Pairing](domain.ReplicaNumber, remote)

	mngr := &usecases.DhtTourDataMngr{
		DhtStr:     dhtStr,
		DhtInfos:   dhtInfos,
		DhtMatches: dhtMatches,
	}
	return mngr
}

// Mock Lider Provider
type MockLeaderProv struct {
}

func (m *MockLeaderProv) GetLeader() string {
	return "192.168.122.219"
}

// Build Worker Client
func BuildWorkerClient(
	elect inter.IElectionPolicy,
	sucProv inter.ISuccProvider) *transport.WorkerTransport {

	addr := "127.0.0.1:" + strconv.Itoa(domain.WClientPort)
	cfg := transport.DefaultCfgAddr(addr)
	// @audit remove
	//leadPr := &MockLeaderProv{}
	// --
	client, err := transport.NewWorkerClient(cfg, elect, sucProv)
	//client, err := transport.NewWorkerClient(cfg, leadPr, sucProv)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Build Worker Mngr
func BuildWorkerMngr() *transport.WorkerMngr {
	addr := ":" + strconv.Itoa(domain.WMngrPort)
	mngr, err := transport.NewWorkerMngr(addr)
	if err != nil {
		log.Fatal("Couldn't intialize worker mngr", err)
	}
	return mngr
}

// Build Middleware
func BuildMiddleware(dm usecases.DataMngr) inter.IMiddleware {
	addr := "localhost:" + strconv.Itoa(domain.MiddPort)
	mid := transport.NewMidServer(addr, dm)
	return mid
}

//func SwapPort(addr string, port uint32) string {
//	addr = addr[:strings.Index(addr, ":")] + ":" + strconv.Itoa(int(port))
//	return addr
//}
