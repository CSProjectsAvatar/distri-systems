package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

type CentDataManager struct {
}

func (dm *CentDataManager) SaveFiles(tour_name string, files map[string]string) error {
	return nil
}

func (dm *CentDataManager) File(tour_name string, file_name string) string {
	return ""
}

func (dm *CentDataManager) SaveMatch(tour_name string, match domain.Pairing) error {
	return nil
}

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
	return "chezz" // @todo implement
}

func (dm *CentDataManager) GetTournInfo(tour_name string) *domain.TournInfo {
	// @todo Mock
	player1 := domain.Player{Name: "Player1"}
	player2 := domain.Player{Name: "Player2"}
	//player3 := domain.Player{Name: "Player3"}

	return &domain.TournInfo{
		Type_:   domain.First_Defeat,
		Players: []domain.Player{player1, player2},
		//Players:  []domain.Player{player1, player2, player3},
	}
}
