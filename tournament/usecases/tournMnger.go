package usecases

import (
	"log"

	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

type TournMngr struct {
	TInfo  *TournInfo
	Winner *Player

	dm               interfaces.DataMngr
	tourTree         *TourNode
	matchsPerPairing map[string]int
}

func (tm *TournMngr) Tree() *TourNode {
	return GetMockTree(tm)
}

func GetMockTree(tm *TournMngr) *TourNode {
	// Mock Tree @todo
	player1 := Player{Id: "Player1"}
	player2 := Player{Id: "Player2"}
	player3 := Player{Id: "Player3"}
	player4 := Player{Id: "Player4"}

	tm.TInfo.Players = []*Player{&player1, &player2, &player3, &player4}

	chP1 := &TourNode{Winner: &player1}
	chP2 := &TourNode{Winner: &player2}
	chP3 := &TourNode{Winner: &player3}
	chP4 := &TourNode{Winner: &player4}

	root := NewNode(tm, DefNodeFunc)

	rch := NewNode(tm, DefNodeFunc)
	rch.SetChildrens([]*TourNode{chP3, chP4})
	root.SetChildrens([]*TourNode{chP1, chP2, rch})

	return root // [p1, p2 [p3, p4]]
}

//func (tm *TournMngr) Tree() *TourNode {
//	switch tm.type_ {
//
//	}
//	return nil
//}

// Returns the name of a Random Unfinished Tournament
func NewRndTour(dm interfaces.DataMngr) *TournMngr {
	tm := &TournMngr{
		dm:               dm,
		matchsPerPairing: make(map[string]int),
	}

	// Initialize the Tournament
	name := dm.UnfinishedTourn()
	tm.TInfo = dm.GetTournInfo(name)
	tm.tourTree = tm.Tree()

	return tm
}

func (tm *TournMngr) GetMatches() <-chan *MatchToRun {
	runnerCh := make(chan *MatchToRun, 10)
	winnerCh := make(chan *Player, 10)

	go func() {
		tm.tourTree.PlayNode(runnerCh, winnerCh)
		tm.Winner = <-winnerCh
		close(runnerCh)
		log.Println("The Winner of the", tm.TInfo.Name, "Tournament is", tm.Winner)
	}() // @audit exception here

	return runnerCh

}

func (tm *TournMngr) AlreadyRun(match *Pairing) (bool, MatchResult) { // @audit implement
	return false, NotPlayed
}

func (tm *TournMngr) GetMatch(pi, pj *Player) *MatchToRun {
	timesPlayed, on_it := tm.matchsPerPairing[pi.Id+pj.Id]
	if on_it {
		tm.matchsPerPairing[pi.Id+pj.Id]++
	} else {
		tm.matchsPerPairing[pi.Id+pj.Id] = 1
	}
	return NewMatchToRun(tm.TInfo.ID, pi, pj, timesPlayed)
}
