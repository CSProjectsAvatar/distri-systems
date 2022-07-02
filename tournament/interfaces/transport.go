package interfaces

import (
	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

type RingTransport interface { // Used on Election
	// Client
	SendToSuccessor(msg *ElectionMsg)
	GetLeaderFromSuccessor() string
	// Server
	MsgNotification() <-chan *ElectionMsg
}

// Responsible for get the matchs for running  them, located in the worker
type IWorkerTransport interface { // Used on Work Client
	GetMatchToRun() *Pairing
	SendResults(match *Pairing) error
}

// Responsible for managing the workers, located in the Leader, server
type IWorkerMngr interface {
	// Gets a worker running the match
	DeliverMatch(match *Pairing)
	// Returns a channel where will be all the results
	NotificationChannel() <-chan *Pairing
}

type ILeaderProvider interface {
	GetLeader() string
}
