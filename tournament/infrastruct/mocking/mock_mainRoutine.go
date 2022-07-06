package mocking

import (
	"log"

	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	tr "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
)

type MockSuccProvider struct {
	succ string
}

func (m *MockSuccProvider) GetSuccessor() string {
	return m.succ
}

func (m *MockSuccProvider) SetSuccessor(succ string) {
	m.succ = succ
}

// --

type MockMainRoutine struct {
	succProvider *MockSuccProvider

	infrastruct.MainRoutine
}

func NewMockRoutine(addr string, succ string) *MockMainRoutine {
	succProv := &MockSuccProvider{succ}
	cfg := tr.DefaultCfgAddr(addr)

	mainR := &infrastruct.MainRoutine{}
	mainR.Elect = infrastruct.NewElectionRingAlgo(addr)           // Election Initialized
	client, err := tr.NewWorkerClient(cfg, mainR.Elect, succProv) // Worker Client Initialized
	if err != nil {
		log.Fatal(err) // @audit Fatal?
	}
	mainR.WClient = client // WClient Setted
	err = mainR.WClient.Start()

	mainR.Elect.SetTransport(client)
	if err != nil {
		log.Fatal(err)
	}

	mock := &MockMainRoutine{
		succProvider: succProv,
		MainRoutine:  *mainR,
	}
	go mainR.WorkDay()
	return mock
}

//set
func (m *MockMainRoutine) SetSuccessor(succ string) {
	m.succProvider.SetSuccessor(succ)
}

//get
func (m *MockMainRoutine) GetSuccessor() string {
	return m.succProvider.GetSuccessor()
}
