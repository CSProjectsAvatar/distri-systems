package usecases

import (
	log "github.com/sirupsen/logrus"

	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

type TournMngr struct {
	TInfo *TournInfo

	dm               DataMngr
	tourTree         *TourNode
	matchsPerPairing map[string]int
	matchsResult     map[string]MatchResult
}

func (tm *TournMngr) Tree() *TourNode {
	SetMockTree(tm)
	return tm.tourTree
}

func (tm *TournMngr) SetTree(tree *TourNode) {
	tm.tourTree = tree
}

func SetMockTree(tm *TournMngr) {
	player1 := Player{Id: "Player1"}
	player2 := Player{Id: "Player2"}
	player3 := Player{Id: "Player3"}
	player4 := Player{Id: "Player4"}

	tm.TInfo.Players = []*Player{&player1, &player2, &player3, &player4}

	chP1 := &TourNode{Winner: &player1}
	chP2 := &TourNode{Winner: &player2}
	chP3 := &TourNode{Winner: &player3}
	chP4 := &TourNode{Winner: &player4}

	root := NewTourNode(tm, DefNodeFunc)

	rch := NewTourNode(tm, DefNodeFunc)
	rch.SetChildrens([]*TourNode{chP3, chP4})
	root.SetChildrens([]*TourNode{chP1, chP2, rch})

	tm.tourTree = root // [p1, p2 [p3, p4]]
}

//func (tm *TournMngr) Tree() *TourNode {
//	switch tm.type_ {
//
//	}
//	return nil
//}

// Returns the name of a Random Unfinished Tournament
func NewRndTour(dm DataMngr) *TournMngr {
	tm := &TournMngr{
		dm:               dm,
		matchsPerPairing: make(map[string]int),
		matchsResult:     make(map[string]MatchResult),
	}

	// Initialize the Tournament
	name, err := dm.UnfinishedTourn()
	if err != nil {
		log.Errorf("Error getting unfinished tournament: %s", err)
		return nil
	}
	tm.TInfo, err = dm.GetTournInfo(name) // @todo check error
	if err != nil {
		log.Errorf("Error getting tournament info: %s", err)
		return nil
	}

	runMatches, err := dm.Matches(tm.TInfo.ID)
	if err != nil {
		log.Errorf("Error getting matches: %s", err)
		return nil
	}
	tm.fillMap(runMatches)

	tm.tourTree = tm.Tree()

	return tm
}

func (tm *TournMngr) GetMatches() <-chan *MatchToRun {
	runnerCh := make(chan *MatchToRun, 10)
	winnerCh := make(chan *Player, 10)

	go func() {
		tm.tourTree.PlayNode(runnerCh, winnerCh)
		tm.TInfo.Winner = <-winnerCh
		close(runnerCh)
		log.Println("The Winner of the", tm.TInfo.Name, "Tournament is", tm.TInfo.Winner.Id)
	}()

	return runnerCh

}

func (tm *TournMngr) AlreadyRun(match *Pairing) (bool, MatchResult) {
	if res, ok := tm.matchsResult[match.ID]; ok {
		return true, res
	}
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

func (tm *TournMngr) fillMap(matches []*Pairing) {
	for _, mtch := range matches {
		tm.matchsResult[mtch.ID] = mtch.Winner
	}
}
