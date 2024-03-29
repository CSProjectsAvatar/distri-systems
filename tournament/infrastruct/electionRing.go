package infrastruct

import (
	. "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	"github.com/CSProjectsAvatar/distri-systems/utils"
	ut "github.com/CSProjectsAvatar/distri-systems/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type ElectionRing struct {
	leader string

	me               string
	coordFlag        bool
	transp           IElectTransport
	notifChn         <-chan *ElectionMsg
	leaderChangedIn  chan<- struct{}
	leaderChangedOut <-chan struct{}

	notNumber int
	ips       []string
}

// MiddlewareIps returns the IPs of nodes under middleware role.
func (ring *ElectionRing) MiddlewareIps() []string {
	//-- @todo llamar eleccion
	return ring.ips
}

func NewElectionRingAlgo(me string) *ElectionRing {
	in, out := utils.MakeInf[struct{}]()
	ring := &ElectionRing{
		me:               me,
		leader:           me, // for working for the one-node ring
		coordFlag:        false,
		leaderChangedIn:  in,
		leaderChangedOut: out,
		ips:              []string{me},
	}
	return ring
}

func (ring *ElectionRing) GetLeader() string {
	return ring.leader
}
func (ring *ElectionRing) GetMe() string {
	return ring.me
}
func (ring *ElectionRing) SetTransport(tr IElectTransport) {
	ring.transp = tr
	ring.notifChn = tr.MsgNotification()

	lead, err := tr.GetLeaderFromSuccessor() // @audit-info you must first enter in chord
	ring.leader = lead
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range ring.notifChn {
			ring.notNumber++
			ring.ElectionMsg(msg)
		}
	}()
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
		ring.leader = ut.Max_in(msg.OnTheRing) // Set leader as the bigger one
		ring.ips = msg.OnTheRing               // saving ring ips so we can send them to client
		ring.leaderChangedIn <- struct{}{}     // Notify that the leader changed
		ring.coordFlag = !ring.coordFlag       // Change flag
		if !ring.coordFlag {                   // Stop? (The flag was Up?)
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

// Returns a Channel that gets a notification when the leader changes
func (ring *ElectionRing) OnLeaderChange() <-chan struct{} {
	return ring.leaderChangedOut
}
