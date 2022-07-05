package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"

	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	inter "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MainRoutine struct {
	DM    usecases.DataMngr
	Elect inter.IElectionPolicy

	ChordSrv *chord.Node
	WClient  inter.IWorkerTransport
	WMngr    inter.IWorkerMngr
	Midd     inter.IMiddleware

	TRunner     inter.Runner
	MatchRunner inter.IMatchRunner
}

func NewMainRoutine(remote *chord.RemoteNode) *MainRoutine {
	m := &MainRoutine{}

	logger := NewLogger().ToFile()
	m.ChordSrv = BuildChordNode(remote, logger) // Chord

	m.DM = BuildDataMngr(m.ChordSrv.Ip, m.ChordSrv.Port) // DataMngr
	sucProv := usecases.NewSuccWrapper(m.ChordSrv)

	addr := m.ChordSrv.Addr()
	cfg := transport.DefaultCfgAddr(addr)

	m.Elect = NewElectionRingAlgo(addr)                // Election
	client := BuildWorkerClient(cfg, m.Elect, sucProv) // WorkerClient

	m.WClient = client // WClient Set
	err := m.WClient.Start()
	if err != nil {
		log.Fatal("Error Starting WorkerClient", err)
	}

	m.Elect.SetTransport(client)

	m.WMngr = BuildWorkerMngr(addr)      // WorkerMngr
	m.Midd = BuildMiddleware(addr, m.DM) // Middleware

	m.TRunner = NewMTRunner(m.WMngr, m.DM) // Runner
	m.MatchRunner = NewWorkerRunner(m.DM)  // MatchRunner

	// Start the Servers
	go m.Midd.Start()

	go m.WorkDay()
	return m
}

func (m *MainRoutine) WorkDay() {
	count := 0
	for {
		match, err := m.WClient.GetMatchToRun() // Subscribe this Worker to Leader
		if err != nil {
			if status.Code(err) != codes.Canceled {
				count++

				if count > domain.MaxRetryTimes {
					m.Elect.CreateElection()
					<-m.Elect.OnLeaderChange() // Wait for Leader Change

					// Run The Mngr if I Am the leader
					if m.IamTheLeader() {
						go m.MngrDay() // Init the Leader Mode
						break
					}

				} else {
					time.Sleep(domain.WhaitTimeBetweenRetry)
				}
			}
		} else {
			match.Winner, err = m.MatchRunner.RunMatch(match)
			if err != nil {
				log.Error(err)
			}

			// Send the Results Back
			err = m.WClient.SendResults(match)
			if err != nil {
				log.Error(err)
			}
		}
	}
	return
}

func (m *MainRoutine) IamTheLeader() bool {
	return m.Elect.GetLeader() == m.Elect.GetMe()
}

// Beguins the control of the flow as a leader
func (m *MainRoutine) MngrDay() {
	runT := make(map[string]bool)
	for {
		// Ask for an Unfinished tournament and Create a TournMngr from it
		tourn, err := usecases.NewRndTour(m.DM)
		if err != nil {
			log.Error("Error Creating a TournMngr")
		} else if !runT[tourn.TInfo.ID] { // if I am not running this tournament
			runT[tourn.TInfo.ID] = true
			go m.TRunner.Run(tourn)
		}
		time.Sleep(5)
		// If I am not longer the leader wet back to work
		if !m.IamTheLeader() {
			go m.WorkDay()
			break
		}
	}
}
