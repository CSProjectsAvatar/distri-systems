package tests

import (
	"math/rand"
	"testing"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	"github.com/stretchr/testify/assert"
)

type MockProvider struct {
	server string
	succ   string
}

func (m *MockProvider) GetLeader() string {
	return m.server
}

func (m *MockProvider) GetSuccessor() string {
	return m.succ
}

func TestWMGiveMeWork_Integration(t *testing.T) {
	assert := assert.New(t)
	serverAddr := "localhost:50051"

	mngr, err := transport.NewWorkerMngr(serverAddr)
	mngr.Start()
	assert.Nil(err)

	confg := transport.DefaultConfig()
	prov := &MockProvider{serverAddr, serverAddr}
	client, err := transport.NewWorkerClient(*confg, prov, prov)
	assert.Nil(err)

	pair1 := &domain.Pairing{ID: "1", TourId: "1", Player1: &domain.Player{Id: "1"}, Player2: &domain.Player{Id: "2"}}
	pair2 := &domain.Pairing{ID: "2", TourId: "1", Player1: &domain.Player{Id: "3"}, Player2: &domain.Player{Id: "4"}}

	go mngr.DeliverMatch(pair1)
	// ask for the first match
	match1, err := client.GetMatchToRun()
	assert.Nil(err)
	assert.Equal(pair1.ID, match1.ID)

	match1.Winner = 1
	// send the result of the first match
	err = client.SendResults(match1)
	assert.Nil(err)

	go mngr.DeliverMatch(pair2)
	// ask for the second match
	match2, err := client.GetMatchToRun()
	assert.Nil(err)
	assert.Equal(pair2.ID, match2.ID)

	match2.Winner = 2
	// send the result of the second match
	err = client.SendResults(match2)
	assert.Nil(err)

	resCh := mngr.NotificationChannel()
	res1 := <-resCh
	res2 := <-resCh
	assert.Equal(res1.ID, match1.ID)
	assert.Equal(res2.ID, match2.ID)
	assert.Equal(domain.MatchResult(match1.Winner), res1.Winner)
	assert.Equal(domain.MatchResult(match2.Winner), res2.Winner)
}

func TestWM_WithMultiplesClients(t *testing.T) {
	assert := assert.New(t)
	serverAddr := "localhost:50051"

	mngr, err := transport.NewWorkerMngr(serverAddr)
	mngr.Start()
	assert.Nil(err)

	confg := transport.DefaultConfig()
	prov := &MockProvider{serverAddr, serverAddr}
	client1, err := transport.NewWorkerClient(*confg, prov, prov)
	assert.Nil(err)
	client2, err := transport.NewWorkerClient(*confg, prov, prov)
	assert.Nil(err)

	pair1 := &domain.Pairing{ID: "1", TourId: "1", Player1: &domain.Player{Id: "1"}, Player2: &domain.Player{Id: "2"}}
	pair2 := &domain.Pairing{ID: "2", TourId: "1", Player1: &domain.Player{Id: "3"}, Player2: &domain.Player{Id: "4"}}
	pair3 := &domain.Pairing{ID: "3", TourId: "1", Player1: &domain.Player{Id: "1"}, Player2: &domain.Player{Id: "4"}}
	pair4 := &domain.Pairing{ID: "4", TourId: "1", Player1: &domain.Player{Id: "3"}, Player2: &domain.Player{Id: "2"}}

	// ask for the pairs to run
	go func() {
		mngr.DeliverMatch(pair1)
		mngr.DeliverMatch(pair2)
		mngr.DeliverMatch(pair3)
		mngr.DeliverMatch(pair4)
	}()

	// ask for the matchs for run in the clients
	go func() {
		for {
			coin := rand.Float32()
			var match *domain.Pairing
			var err error

			if coin < 0.5 {
				match, err = client1.GetMatchToRun()
			} else {
				match, err = client2.GetMatchToRun()
			}
			assert.Nil(err)
			assert.NotNil(match)

			match.Winner = 1
			// send the result of the match
			err = client1.SendResults(match)
			assert.Nil(err)
		}
	}()

	outCh := make(chan struct{}, 1)
	// Review Results
	go func() {
		ch := mngr.NotificationChannel()
		var matchList []*domain.Pairing

		for len(matchList) < 4 {
			matchList = append(matchList, <-ch)
		}
		assert.Equal(4, len(matchList))
		originalPairs := []*domain.Pairing{pair1, pair2, pair3, pair4}
		for _, pair := range originalPairs {
			assert.True(ContainsPairAndNotNilWinner(matchList, pair))
		}
		outCh <- struct{}{}
	}()
	<-outCh
}

func ContainsPairAndNotNilWinner(list []*domain.Pairing, pair *domain.Pairing) bool {
	for _, p := range list {
		if p.Winner == 0 {
			break
		}
		if p.ID == pair.ID {
			return true
		}
	}
	return false
}
