package interfaces

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

type RingTransport interface { // Used on Election
	SendToSuccessor(msg *ElectionMsg)
	GetLeaderFromSuccessor() string
}

type WorkerTransport interface { // Used on Work Client
	SendResults(match domain.Pairing) error
	SubscribeForWork() error
}
