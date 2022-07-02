package transport

import (
	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	pb "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_workerMngr"
)

func NewResultReq(match *Pairing) *pb.ResultReq {
	return &pb.ResultReq{
		MatchId:     match.ID,
		TourId:      match.TourId,
		FstPlayerID: match.Player1.Id,
		SndPlayerID: match.Player2.Id,
		Winner:      uint32(match.Winner),
	}
}

func FromMatchResp(match *pb.MatchResp) *Pairing {
	return &Pairing{
		ID:      match.MatchId,
		TourId:  match.TourId,
		Player1: &Player{Id: match.FstPlayerID},
		Player2: &Player{Id: match.SndPlayerID},
	}
}
