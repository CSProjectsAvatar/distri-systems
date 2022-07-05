package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	inter "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// Build Chord Node
func BuildChordNode(ip string, remote *chord.RemoteNode, logger domain.Logger) *chord.Node {
	entry, err := usecases.NewNode(
		ChordConfig(ip, domain.ChordPort),
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

// Build Worker Client
func BuildWorkerClient(
	cfg *transport.Config,
	elect inter.IElectionPolicy,
	sucProv inter.ISuccProvider) *transport.WorkerTransport {

	client, err := transport.NewWorkerClient(*cfg, elect, sucProv)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Build Worker Mngr
func BuildWorkerMngr(addr string) *transport.WorkerMngr {
	// swap port by the one on domain
	addr = SwapPort(addr, domain.WMngrPort)

	mngr, err := transport.NewWorkerMngr(addr)
	if err != nil {
		log.Fatal("Couldn't intialize worker mngr", err)
	}
	return mngr
}

// Build Middleware
func BuildMiddleware(addr string, dm usecases.DataMngr) inter.IMiddleware {
	addr = SwapPort(addr, domain.MiddPort)
	mid := transport.NewMidServer(addr, dm)
	return mid
}

func SwapPort(addr string, port uint32) string {
	addr = addr[:strings.Index(addr, ":")] + ":" + strconv.Itoa(int(port))
	return addr
}
