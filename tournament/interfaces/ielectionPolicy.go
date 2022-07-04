package interfaces

// Policy that assure always a Leader
type IElectionPolicy interface {
	CreateElection()
	ElectionMsg(msg *ElectionMsg)
	OnLeaderChange() <-chan struct{}
	GetLeader() string
	GetMe() string
	SetTransport(transport IElectTransport)
}

type ElectionType int

const (
	ELECTION ElectionType = iota
	COORDINATOR
)

type ElectionMsg struct {
	Type      ElectionType
	OnTheRing []string
}
