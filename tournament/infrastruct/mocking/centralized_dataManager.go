package mocking

import (
	"math/rand"
	"strconv"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/utils"
)

type CentDataManager struct {
}

func (dm *CentDataManager) SaveFiles(tour_name string, files *map[string]string) error {
	return nil
}

func (dm *CentDataManager) File(tour_name string, file_name string) string {
	return ""
}

func (dm *CentDataManager) SaveMatch(match *domain.Pairing) error {
	return nil
}

// func (dm *CentDataManager) SaveMatch(tour_name string, match *domain.Pairing) error {
// 	return nil
// }

func (dm *CentDataManager) Matches(tour_name string) []*domain.Pairing {
	return nil
}

func (dm *CentDataManager) SaveTournTree(tour_name string, tree *domain.TourNode) error {
	return nil
}

//func (dm *CentDataManager) TournTree(tour_name string) *domain.TourNode {
//	// Mock Tree @todo
//	return GetMockTree()
//}

func (dm *CentDataManager) UnfinishedTourn() string {
	random := rand.Int()
	return "tour_" + strconv.Itoa(random)
}

func (dm *CentDataManager) GetTournInfo(tour_name string) *domain.TournInfo {
	// @todo Mock
	player1 := domain.Player{Id: "Player1"}
	player2 := domain.Player{Id: "Player2"}
	//player3 := domain.Player{Id: "Player3"}

	return &domain.TournInfo{
		Name: "tour_" + tour_name,
		Players: []*domain.Player{
			&player1,
			&player2,
			//&player3,
		},
		Type_: domain.All_vs_All,
		ID:    utils.Hash(tour_name + string(domain.All_vs_All)),
	}
}
