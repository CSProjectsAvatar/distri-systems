package infrastruct

import (
	inter "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	ut "github.com/CSProjectsAvatar/distri-systems/utils"
	"golang.org/x/exp/slices"
)

type ElectionRing struct {
	Leader string

	me        string
	coordFlag bool
	transp    inter.RingTransport
}

func NewElectionRingAlgo(me string, transp inter.RingTransport) *ElectionRing {
	return &ElectionRing{
		me:        me,
		coordFlag: false,
		transp:    transp,
		Leader:    transp.GetLeaderFromSuccessor(), // @audit-info you must first enter in chord
	}
}

func (ring *ElectionRing) ElectionMsg(msg *inter.ElectionMsg) {
	switch msg.Type {
	case inter.ELECTION:
		ring.coordFlag = false                                  // Flag v
		stop := slices.Contains[string](msg.OnTheRing, ring.me) // Me on List
		if stop {
			ring.coordFlag = true        // Flag ^
			msg.Type = inter.COORDINATOR // Change msg.Type
		} else {
			msg.OnTheRing = append(msg.OnTheRing, ring.me) // Add me to list
		}

	case inter.COORDINATOR:
		ring.Leader = ut.Max_in(msg.OnTheRing)  // Set leader as the bigger one
		ring.coordFlag = ring.coordFlag != true // Change flag
		if !ring.coordFlag {                    // Stop? (The flag was Up?)
			return
		}
	}
	ring.transp.SendToSuccessor(msg) // Send msg to successor
}

func (ring *ElectionRing) CreateElection() {
	msg := &inter.ElectionMsg{
		Type:      inter.ELECTION,
		OnTheRing: []string{ring.me},
	}
	ring.transp.SendToSuccessor(msg) // Send msg to successor
}
