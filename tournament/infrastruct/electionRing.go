package infrastruct

import (
	. "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	ut "github.com/CSProjectsAvatar/distri-systems/utils"
	"golang.org/x/exp/slices"
)

type ElectionRing struct {
	Leader string

	me        string
	coordFlag bool
	transp    RingTransport
	notifChn  <-chan *ElectionMsg
}

func NewElectionRingAlgo(me string, transp RingTransport) *ElectionRing {
	ring := &ElectionRing{
		me:        me,
		coordFlag: false,
		transp:    transp,
		notifChn:  transp.MsgNotification(),
		Leader:    transp.GetLeaderFromSuccessor(), // @audit-info you must first enter in chord
	}
	go func() {
		for msg := range ring.notifChn {
			ring.ElectionMsg(msg)
		}
	}()
	return ring
}

func (ring *ElectionRing) ElectionMsg(msg *ElectionMsg) {
	switch msg.Type {
	case ELECTION:
		ring.coordFlag = false                                  // Flag v
		stop := slices.Contains[string](msg.OnTheRing, ring.me) // Me on List
		if stop {
			ring.coordFlag = true  // Flag ^
			msg.Type = COORDINATOR // Change msg.Type
		} else {
			msg.OnTheRing = append(msg.OnTheRing, ring.me) // Add me to list
		}

	case COORDINATOR:
		ring.Leader = ut.Max_in(msg.OnTheRing)  // Set leader as the bigger one
		ring.coordFlag = ring.coordFlag != true // Change flag
		if !ring.coordFlag {                    // Stop? (The flag was Up?)
			return
		}
	}
	ring.transp.SendToSuccessor(msg) // Send msg to successor
}

func (ring *ElectionRing) CreateElection() {
	msg := &ElectionMsg{
		Type:      ELECTION,
		OnTheRing: []string{ring.me},
	}
	ring.transp.SendToSuccessor(msg) // Send msg to successor
}
