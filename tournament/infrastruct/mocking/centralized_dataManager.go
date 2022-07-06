package mocking

import (
	"math/rand"
	"strconv"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/utils"
)

type CentDataManager struct {
}

func (dm *CentDataManager) SetTournInfo(info *domain.TournInfo) error {
	//TODO implement me
	return nil
}

func (dm *CentDataManager) SaveFiles(tour_name string, files *map[string]string) error {
	return nil
}

func (dm *CentDataManager) File(tourId string, fileName string) (string, error) {
	return "", nil
}

func (dm *CentDataManager) FileGroup(tourId string, files []string) (map[string]string, error) {
	return nil, nil
}
func (dm *CentDataManager) SaveMatch(match *domain.Pairing) error {
	return nil
}

// func (dm *CentDataManager) SaveMatch(tour_name string, match *domain.Pairing) error {
// 	return nil
// }

func (dm *CentDataManager) Matches(tourId string) ([]*domain.Pairing, error) {
	return nil, nil
}

func (dm *CentDataManager) SaveTournTree(tour_name string, tree *domain.TourNode) error {
	return nil
}

//func (dm *CentDataManager) TournTree(tour_name string) *domain.TourNode {
//	// Mock Tree @todo
//	return GetMockTree()
//}

func (dm *CentDataManager) UnfinishedTourn() (string, error) {
	random := rand.Int()
	return "tour_" + strconv.Itoa(random), nil
}

func (dm *CentDataManager) GetTournInfo(tourId string) (*domain.TournInfo, error) {
	// @todo Mock
	player1 := domain.Player{Id: "Player1"}
	player2 := domain.Player{Id: "Player2"}
	//player3 := domain.Player{Id: "Player3"}

	return &domain.TournInfo{
		Name: "tour_" + tourId,
		Players: []*domain.Player{
			&player1,
			&player2,
			//&player3,
		},
		Type_: domain.All_vs_All,
		ID:    utils.Hash(tourId + string(domain.All_vs_All)),
	}, nil
}

//GetAllIds
func (dm *CentDataManager) GetAllIds() ([]string, error) {
	return nil, nil
}
