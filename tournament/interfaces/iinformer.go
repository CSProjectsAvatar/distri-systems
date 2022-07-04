package interfaces

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

type IInformer interface {
	// GetStatistics get statistics for a given tournament
	GetStatistics(tourName string) (*domain.TournInfo, []*domain.Pairing)
}
