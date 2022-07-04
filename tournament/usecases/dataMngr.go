package usecases

import (
	"errors"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

// DataMngr is responsible for Save and Retrieve Data of tournaments.
type DataMngr interface {
	// SaveFiles Saves the tournament files. 'files' is a map of file names to file contents.
	SaveFiles(tourId string, files *map[string]string) error

	// File Retrieves a file from the tournament.
	File(tourId string, fileName string) (string, error) // @todo now an error is returned, consider that in callers

	// SaveMatch Saves a match already run
	SaveMatch(match *domain.Pairing) error

	// Matches Retrieves the tournament's matches
	Matches(tourId string) ([]*domain.Pairing, error) // @todo now an error is returned, consider that in callers

	// GetTournInfo Loads a tournament main info. Returns ErrInfoNotfound if... (u know).
	GetTournInfo(tourId string) (*domain.TournInfo, error) // @todo now an error is returned, consider that in callers

	SetTournInfo(info *domain.TournInfo) error

	// UnfinishedTourn Loads a tournament not finished. Returns its ID.
	UnfinishedTourn() (string, error)
}

var ErrInfoNotFound = errors.New("info not found")