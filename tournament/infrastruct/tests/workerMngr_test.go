package tests

import (
	"context"
	"testing"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	pb "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_workerMngr"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/transport"
	"github.com/stretchr/testify/assert"
)

func TestWMGiveMeWork(t *testing.T) {
	assert := assert.New(t)
	addr := "localhost:50051"
	mngr, err := transport.NewWorkerMngr(addr)
	if err != nil {
		assert.Nil(err)
	}
	pair1 := &domain.Pairing{ID: "1", TourId: "1", Player1: &domain.Player{Id: "1"}, Player2: &domain.Player{Id: "2"}}
	pair2 := &domain.Pairing{ID: "2", TourId: "1", Player1: &domain.Player{Id: "3"}, Player2: &domain.Player{Id: "4"}}

	go mngr.DeliverMatch(pair1)
	resp, err := mngr.GiveMeWork(context.Background(), &pb.MatchReq{})
	assert.Nil(err)
	assert.Equal(pair1.ID, resp.MatchId)

	go mngr.DeliverMatch(pair2)
	resp, err = mngr.GiveMeWork(context.Background(), &pb.MatchReq{})
	assert.Nil(err)
	assert.Equal(pair2.ID, resp.MatchId)
}

func TestWMCatchReq(t *testing.T) {
	assert := assert.New(t)
	addr := "localhost:50052"
	mngr, err := transport.NewWorkerMngr(addr)
	if err != nil {
		assert.Nil(err)
	}
	pair1 := &domain.Pairing{ID: "1", TourId: "1", Player1: &domain.Player{Id: "1"}, Player2: &domain.Player{Id: "2"}}
	pair2 := &domain.Pairing{ID: "2", TourId: "1", Player1: &domain.Player{Id: "3"}, Player2: &domain.Player{Id: "4"}}

	// Send the first match to run
	go mngr.DeliverMatch(pair1)
	resp, err := mngr.GiveMeWork(context.Background(), &pb.MatchReq{})
	assert.Nil(err)
	assert.Equal(pair1.ID, resp.MatchId)

	// Send the second match to run
	go mngr.DeliverMatch(pair2)
	resp2, err := mngr.GiveMeWork(context.Background(), &pb.MatchReq{})
	assert.Nil(err)
	assert.Equal(pair2.ID, resp2.MatchId)

	// Simulate the result of the first and second matchs
	resReq := &pb.ResultReq{
		MatchId:     resp.MatchId,
		TourId:      resp.TourId,
		FstPlayerID: resp.FstPlayerID,
		SndPlayerID: resp.SndPlayerID,
		Winner:      1,
	}
	resReq2 := &pb.ResultReq{
		MatchId:     resp2.MatchId,
		TourId:      resp2.TourId,
		FstPlayerID: resp2.FstPlayerID,
		SndPlayerID: resp2.SndPlayerID,
		Winner:      2,
	}
	//--
	// Send the result
	mngr.CatchResult(context.Background(), resReq)
	mngr.CatchResult(context.Background(), resReq2)

	notChn := mngr.NotificationChannel()
	res1 := <-notChn
	res2 := <-notChn

	assert.Equal(resReq.MatchId, res1.ID)
	assert.Equal(domain.MatchResult(resReq.Winner), res1.Winner)

	assert.Equal(resReq2.MatchId, res2.ID)
	assert.Equal(domain.MatchResult(resReq2.Winner), res2.Winner)
}
