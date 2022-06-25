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
	tm.name = dm.UnfinishedTourn() // @todo When exist 2 leaders, manage syncronization
	ti := dm.GetTournInfo(tm.name)
	tm.type_, tm.players = ti.Type_, ti.Players
	tm.tourTree = tm.Tree()
	return tm
}

func (tm *TournMngr) GetMatches() <-chan *domain.MatchToRun {
	runnerCh := make(chan *domain.MatchToRun, 10)
	winnerCh := make(chan *domain.Player, 10)

	go func() {
		tm.tourTree.PlayNode(runnerCh, winnerCh)
		tm.winner = <-winnerCh
		fmt.Println("The Winner of the Tournament is", tm.winner)
		// @todo call dm here
	}() // @audit exception here

	return runnerCh
}
