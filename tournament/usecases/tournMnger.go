package usecases

import (
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

type TournMngr struct {
	dm interfaces.DataMngr

	name     string
	type_    domain.TourType
	players  []domain.Player
	tourTree *domain.TourNode
	winner   *domain.Player
}

func (tm *TournMngr) Tree() *domain.TourNode {
	switch tm.type_ {

	}
	return nil
}

// Returns the name of a Random Unfinished Tournament
func NewRndTour(dm interfaces.DataMngr) *TournMngr {
	tm := new(TournMngr)

	// Initialize the Tournament
	tm.name = dm.RndUnfinishedTourn()
	ti := dm.GetTournInfo(tm.name)
	tm.type_, tm.players, tm.tourTree = ti.Type_, ti.Players, ti.TourTree

	if tm.tourTree == nil { // If there is no tour tree, create one
		tm.Tree()
		dm.SaveTournTree(tm.name, tm.tourTree) // Save the tour tree
	}
	return tm
}

func (tm *TournMngr) GetMatches() <-chan *domain.MatchToRun {
	runnerCh := make(chan *domain.MatchToRun, 10)
	winnerCh := make(chan *domain.Player, 10)

	go func() {
		tm.tourTree.PlayNode(runnerCh, winnerCh)
		tm.winner = <-winnerCh
		fmt.Println("The Winner of the Tournament is", tm.winner)
	}() // @audit exception here

	return runnerCh
}
