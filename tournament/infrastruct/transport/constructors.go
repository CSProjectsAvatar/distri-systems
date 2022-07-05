package transport

import (
	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	pb_m "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_mid"
	pb_r "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_ring"
	pb_w "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_workerMngr"
	. "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

func NewResultReq(match *Pairing) *pb_w.ResultReq {
	return &pb_w.ResultReq{
		MatchId:     match.ID,
		TourId:      match.TourId,
		FstPlayerID: match.Player1.Id,
		SndPlayerID: match.Player2.Id,
		Winner:      uint32(match.Winner),
	}
}

func NewElectionMsgReq(msg *ElectionMsg) *pb_r.ElectionMsgReq {
	return &pb_r.ElectionMsgReq{
		Type: uint32(msg.Type),
		OnIt: msg.OnTheRing,
	}

}

func FromMatchResp(match *pb_w.MatchResp) *Pairing {
	return &Pairing{
		ID:      match.MatchId,
		TourId:  match.TourId,
		Player1: &Player{Id: match.FstPlayerID},
		Player2: &Player{Id: match.SndPlayerID},
	}
}

func ToTournFiles(files []*pb_m.File) []*TournFile {
	var tournFiles []*TournFile
	for _, file := range files {
		tournFiles = append(tournFiles, &TournFile{
			Name:   file.Name,
			Data:   file.Data,
			IsGame: file.IsGame,
		})
	}
	return tournFiles
}
