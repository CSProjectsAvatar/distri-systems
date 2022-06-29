package interfaces

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

// Responsible for managing the workers, located in the Leader
type IWorkerMngr interface {
	DeliverMatch(match *domain.Pairing)
	NotificationChannel() chan *domain.Pairing
}
