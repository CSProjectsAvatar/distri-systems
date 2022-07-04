package usecases

import (
	// . "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

	// infra "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	// use "github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	// tr "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	inter "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MainRoutine struct {
	Elect   inter.IElectionPolicy
	WClient inter.IWorkerTransport
	WMngr   inter.IWorkerMngr
	DM      inter.DataMngr
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
					// @audit Here Goes The Code to Run The Mngr if I Am the leader
					if m.Elect.GetLeader() == m.Elect.GetMe() {
						go m.WMngr.Start()
						break
					}

				} else {
					time.Sleep(domain.WhaitTimeBetweenRetry)
				}
			}
		} else {
			// @audit Mock Match Result for now
			match.Winner = 2

			// Send the Results Back
			err = m.WClient.SendResults(match)
			if err != nil {
				log.Error(err)
			}
		}
	}
	return
}

// Beguins the control of the flow as a leader
// func (m *MainRoutine) MngrDay() {
// 	// Ask for an Unfinished tournament and Create a TournMngr from it
// 	tourID, err := m.DM.UnfinishedTourn()
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	tInfo := m.DM.GetTournInfo(tourID)
// 	tMngr := domain.NewTournMngr(tInfo)
// 	// Initialize a Runner
// 	// Run the tournaments
// }

// func NewMainRoutine(addr string) *MainRoutine {

// cfg := tr.DefaultCfgAddr(addr)

// mainR := &use.MainRoutine{}
// mainR.Elect = infra.NewElectionRingAlgo(addr)           // Election Initialized
// client, err := tr.NewWorkerClient(cfg, mainR.Elect, succProv) // Worker Client Initialized
// if err != nil {
// 	log.Fatal(err) // @audit Fatal?
// }
// mainR.WClient = client // WClient Setted
// err = mainR.WClient.Start()

// mainR.Elect.SetTransport(client)
// if err != nil {
// 	log.Fatal(err)
// }
// 	elect := NewElectionRingAlgo(addr)
// 	wClient := NewWorkerClient(addr, elect, succProv)

// }
