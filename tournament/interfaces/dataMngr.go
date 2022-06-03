package interfaces

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

// Responsible for Save and Retrieve Data
type DataMngr interface {
	SaveFiles(tour_name string, files map[string]string) error
	File(tour_name string, file_name string) string

	SaveStat(tour_name string, match domain.Match) error
	Stats(tour_name string) []domain.Match

	SaveTournTree(tour_name string, tree *domain.TourNode) error
	TournTree(tour_name string) *domain.TourNode

	RndUnfinishedTourn() string

	GetTournInfo(tour_name string) *domain.TournInfo
}
