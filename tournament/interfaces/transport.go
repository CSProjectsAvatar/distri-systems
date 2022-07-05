package interfaces

import (
	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

type ITransportSrvr interface {
	Start() error
	Stop() error
}

type IElectTransport interface { // Used on Election
	// Client
	SendToSuccessor(msg *ElectionMsg) error
	GetLeaderFromSuccessor() (string, error)
	// Server
	MsgNotification() <-chan *ElectionMsg
}

// Responsible for get the matchs for running  them, located in the worker
type IWorkerTransport interface { // Used on Work Client
	ITransportSrvr
	GetMatchToRun() (*Pairing, error)
	SendResults(match *Pairing) error
}

// Responsible for managing the workers, located in the Leader, server
type IWorkerMngr interface {
	ITransportSrvr
	// Gets a worker running the match
	DeliverMatch(match *Pairing)
	// Returns a channel where will be all the results
	NotificationChannel() <-chan *Pairing
}

type IMiddleware interface {
	ITransportSrvr
}

type ILeaderProvider interface {
	GetLeader() string
}

type ISuccProvider interface {
	GetSuccessor() string
}
