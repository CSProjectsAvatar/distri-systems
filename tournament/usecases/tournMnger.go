package usecases

import (
	"log"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

type TournMngr struct {
	TInfo  *domain.TournInfo
	Winner *domain.Player

	dm       interfaces.DataMngr
	tourTree *domain.TourNode
}

func (tm *TournMngr) Tree() *domain.TourNode {
	return GetMockTree()
}

func GetMockTree() *domain.TourNode {
	// Mock Tree @todo
	player1 := domain.Player{Name: "Player1"}
	player2 := domain.Player{Name: "Player2"}
	player3 := domain.Player{Name: "Player3"}
	player4 := domain.Player{Name: "Player4"}

	return &domain.TourNode{
		Children: []*domain.TourNode{
			&domain.TourNode{
				Winner: &player1,
			},
			&domain.TourNode{
				Winner: &player2,
			},
			&domain.TourNode{
				Children: []*domain.TourNode{
					&domain.TourNode{
						Winner: &player3,
					},
					&domain.TourNode{
						Winner: &player4,
					},
				},
				JoinFunc: domain.DefNodeFunc,
			},
		},
		JoinFunc: domain.DefNodeFunc,
	}
}

//func (tm *TournMngr) Tree() *domain.TourNode {
//	switch tm.type_ {
//
//	}
//	return nil
//}

// Returns the name of a Random Unfinished Tournament
func NewRndTour(dm interfaces.DataMngr) *TournMngr {
	tm := &TournMngr{dm: dm}

	// Initialize the Tournament
	name := dm.UnfinishedTourn()
	tm.TInfo = dm.GetTournInfo(name)
	tm.tourTree = tm.Tree()

	return tm
}

func (tm *TournMngr) GetMatches() <-chan *domain.MatchToRun {
	runnerCh := make(chan *domain.MatchToRun, 10)
	winnerCh := make(chan *domain.Player, 10)

	go func() {
		tm.tourTree.PlayNode(runnerCh, winnerCh)
		tm.Winner = <-winnerCh
		close(runnerCh)
		log.Println("The Winner of the", tm.TInfo.Name, "Tournament is", tm.Winner)
	}() // @audit exception here

	return runnerCh

}

func (tm *TournMngr) AlreadyRun(match *domain.Pairing) (bool, domain.MatchResult) {
	return false, domain.NotPlayed
}
