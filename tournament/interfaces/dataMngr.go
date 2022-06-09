package interfaces

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

// Responsible for Save and Retrieve Data
type DataMngr interface {
	// SaveFiles Saves the tournament files. 'files' is a map of file names to file contents.
	SaveFiles(tour_name string, files map[string]string) error

	// File Retrieves a file from the tournament.
	File(tour_name string, file_name string) string

	// SaveMatch Saves a match already run
	SaveMatch(tour_name string, match domain.Match) error

	// Matches Retrieves the tournament's matches
	Matches(tour_name string) []*domain.Match

	// GetTournInfo Loads a tournament main info
	GetTournInfo(tour_name string) *domain.TournInfo

	// UnfinishedTourn Loads a tournament not finished
	UnfinishedTourn() string
}
