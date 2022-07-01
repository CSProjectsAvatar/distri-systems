package interfaces

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

// Responsible for managing the workers, located in the Leader
type IWorkerMngr interface {
	// Gets a worker running the match
	DeliverMatch(match *domain.Pairing)
	// Returns a channel where will be all the results
	NotificationChannel() chan *domain.Pairing
}
