package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/utils"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"

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
	switch tm.TInfo.Type_ {
	case First_Defeat:
		return tm.BuildFirstDefeat()
	case All_vs_All:
		return tm.BuildAllVsAll()
	}
	return nil
}

func (tm *TournMngr) BuildFirstDefeat() *TourNode {
	var tourNodes []*TourNode

	for i := 0; i < len(tm.TInfo.Players); i++ {
		tourNodes = append(tourNodes, &TourNode{Children: nil, Winner: tm.TInfo.Players[i]})
	}

	for len(tourNodes) > 1 {
		for j := 0; j < len(tourNodes); j += 2 {
			right := tourNodes[j]
			left := tourNodes[j+1]
			children := []*TourNode{left, right}
			tourNode := &TourNode{Children: children, Winner: &Player{}}
			var tourNodes1 = append(tourNodes[0:j], tourNode)
			tourNodes = append(tourNodes1, tourNodes[j+2:]...)
		}
	}

	return tourNodes[0]
}

func (tm *TournMngr) BuildAllVsAll() *TourNode {
	var children []*TourNode
	var root *TourNode = &TourNode{Children: children, Winner: &Player{}}
	for i := 0; i < len(tm.TInfo.Players); i++ {
		var child *TourNode = &TourNode{Children: nil, Winner: tm.TInfo.Players[i]}
		root.Children = append(root.Children, child)
	}
	return root
}

// func (tm *TournMngr) Tree() *TourNode {
// 	SetMockTree(tm)
// 	return tm.tourTree
// }

func (tm *TournMngr) SetTree(tree *TourNode) {
	tm.tourTree = tree
	tm.tourTree.SetJoinFunc(DefNodeFunc)
	tm.tourTree.SetProvider(tm)
}

func NewMockTourMngr() *TournMngr {
	// dm := &mocking.CentDataManager{}
	tm := &TournMngr{
		// dm:               dm,
		matchsPerPairing: make(map[string]int),
		matchsResult:     make(map[string]MatchResult),
		TInfo: &TournInfo{
			ID:      utils.Hash("MockTour1" + time.Now().String()),
			Name:    "MockTour" + strconv.Itoa(rand.Int())[:2],
			Type_:   First_Defeat,
			Players: []*Player{&Player{Id: "1"}, &Player{Id: "2"}, &Player{Id: "3"}, &Player{Id: "4"}},
		},
	}
	SetMockTree(tm)
	return tm
}
func SetMockTree(tm *TournMngr) {
	player1 := Player{Id: "Player1"}
	player2 := Player{Id: "Player2"}
	player3 := Player{Id: "Player3"}
	player4 := Player{Id: "Player4"}

	tm.TInfo.Players = []*Player{&player1, &player2, &player3, &player4}
	for _, player := range tm.TInfo.Players {
		player.Id += "_" + strconv.Itoa(rand.Int())[:3]
	}

	chP1 := &TourNode{Winner: &player1}
	chP2 := &TourNode{Winner: &player2}
	chP3 := &TourNode{Winner: &player3}
	chP4 := &TourNode{Winner: &player4}

	root := NewTourNode(tm, DefNodeFunc)

	rch := NewTourNode(tm, DefNodeFunc)
	rch.SetChildrens([]*TourNode{chP3, chP4})
	root.SetChildrens([]*TourNode{chP1, chP2, rch})

	root.SetProvider(tm)
	tm.SetTree(root) // [p1, p2, [p3, p4]]
}

// Returns the name of a Random Unfinished Tournament
func NewRndTour(dm DataMngr) (*TournMngr, error) {
	tm := &TournMngr{
		dm:               dm,
		matchsPerPairing: make(map[string]int),
		matchsResult:     make(map[string]MatchResult),
	}

	// Initialize the Tournament
	name, err := dm.UnfinishedTourn()
	if err != nil {
		log.Errorf("Error getting unfinished tournament: %s", err)
		return nil, err
	}

	tm.TInfo, err = dm.GetTournInfo(name) // @todo check error
	if err != nil {
		log.Errorf("Error getting tournament info: %s", err)
		return nil, err
	}

	runMatches, err := dm.Matches(tm.TInfo.ID)
	if err != nil {
		log.Errorf("Error getting matches: %s", err)
		// return nil // @audit Not Critical Error
	}
	tm.fillMap(runMatches)

	tm.SetTree(tm.Tree())

	return tm, nil
}

func (tm *TournMngr) GetMatches() <-chan *MatchToRun {
	runnerCh := make(chan *MatchToRun, 10)
	winnerCh := make(chan *Player, 10)

	go func() {
		tm.tourTree.PlayNode(runnerCh, winnerCh)
		tm.TInfo.Winner = <-winnerCh
		close(runnerCh)
		// @audit save winner in db

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
	if matches == nil {
		return
	}
	for _, mtch := range matches {
		tm.matchsResult[mtch.ID] = mtch.Winner
	}
}
