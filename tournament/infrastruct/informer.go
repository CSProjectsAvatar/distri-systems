package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/mocking"
)

type Informer struct {
}

func (info *Informer) GetStatistics(tourName string) (*domain.TournInfo, []*domain.Pairing) {
	manager := &mocking.CentDataManager{}
	tournInfo := manager.GetTournInfo(tourName)
	matches := manager.Matches(tourName)

	return tournInfo, matches
}
