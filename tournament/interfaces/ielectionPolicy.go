package interfaces

// Policy that assure always a Leader
type ElectionPolicy interface {
	ElectionMsg(msg *ElectionMsg)
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
