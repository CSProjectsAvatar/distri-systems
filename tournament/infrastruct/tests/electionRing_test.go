package test

import (
	// "log"

	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	inf "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

// func Init() {
// 	return electionR
// }
func TestElectionArrive_NotInList(t *testing.T) {
	log.Println("> Election - Not in List <")

	transp := &MockTransporter{}
	electionR := inf.NewElectionRingAlgo("nodo1")
	electionR.SetTransport(transp)

	electMsg := &interfaces.ElectionMsg{
		Type:      interfaces.ELECTION,
		OnTheRing: []string{"nodo2"},
	}
	//--
	electionR.ElectionMsg(electMsg)

	out := transp.lastMsgToSuccessor
	assert.Equal(t, interfaces.ELECTION, out.Type)
	assert.Equal(t, []string{"nodo2", "nodo1"}, out.OnTheRing)
}

func TestElectionArrive_InList(t *testing.T) {
	log.Println("> Election - In List <")

	transp := &MockTransporter{}
	electionR := inf.NewElectionRingAlgo("nodo1")
	electionR.SetTransport(transp)

	electMsg := &interfaces.ElectionMsg{
		Type:      interfaces.ELECTION,
		OnTheRing: []string{"nodo1", "nodo2"},
	}
	//--
	electionR.ElectionMsg(electMsg)

	out := transp.lastMsgToSuccessor
	assert.Equal(t, interfaces.COORDINATOR, out.Type)
	assert.Equal(t, []string{"nodo1", "nodo2"}, out.OnTheRing)
}

func TestCoordinatorArrive_FirstTime(t *testing.T) {
	log.Println("> Coordinator - First Time <")

	transp := &MockTransporter{}
	electionR := inf.NewElectionRingAlgo("nodo1")
	electionR.SetTransport(transp)

	electMsg := &interfaces.ElectionMsg{
		Type:      interfaces.COORDINATOR,
		OnTheRing: []string{"nodo1", "nodo2", "nodo3"},
	}
	//--
	electionR.ElectionMsg(electMsg)

	out := transp.lastMsgToSuccessor
	assert.Equal(t, interfaces.COORDINATOR, out.Type)
	assert.Equal(t, "nodo3", electionR.GetLeader())
}

func TestCoordinatorArrive_SecondTime(t *testing.T) {
	log.Println("> Coordinator - Second Time <")

	transp := &MockTransporter{}
	electionR := inf.NewElectionRingAlgo("nodo1")
	electionR.SetTransport(transp)

	electMsg := &interfaces.ElectionMsg{
		Type:      interfaces.COORDINATOR,
		OnTheRing: []string{"nodo1", "nodo2", "nodo3"},
	}
	electMsg2 := &interfaces.ElectionMsg{
		Type:      interfaces.COORDINATOR,
		OnTheRing: []string{"nodo1", "nodo4", "nodo3"},
	}
	//--
	electionR.ElectionMsg(electMsg)
	electionR.ElectionMsg(electMsg2)

	out := transp.lastMsgToSuccessor
	assert.Equal(t, interfaces.COORDINATOR, out.Type)
	assert.Equal(t, "nodo4", electionR.GetLeader())
	assert.Equal(t, electMsg, out)
}

func Test_Election_Then_Coordinator(t *testing.T) {
	log.Println("> Election - Then - Coordinator <")

	transp := &MockTransporter{}
	electionR := inf.NewElectionRingAlgo("nodo1")
	electionR.SetTransport(transp)

	electMsg := &interfaces.ElectionMsg{
		Type:      interfaces.ELECTION,
		OnTheRing: []string{"nodo1", "nodo2", "nodo3"},
	}
	//--
	electionR.ElectionMsg(electMsg)            // election arrives
	assert.Equal(t, "", electionR.GetLeader()) // no GetLeader() yet
	coord := transp.lastMsgToSuccessor         // coord msg out
	transp.lastMsgToSuccessor = nil            // reset
	electionR.ElectionMsg(coord)               // coord arrives, no other msg should be sent

	out := transp.lastMsgToSuccessor
	assert.Equal(t, interfaces.COORDINATOR, coord.Type)
	assert.Equal(t, "nodo3", electionR.GetLeader())
	assert.Nil(t, out)
}

type MockTransporter struct {
	lastMsgToSuccessor *interfaces.ElectionMsg
	actualLeader       string
}

func (m *MockTransporter) SendToSuccessor(msg *interfaces.ElectionMsg) error {
	log.Println("SendToSuccessor:", msg)
	m.lastMsgToSuccessor = msg
	return nil
}

func (m *MockTransporter) GetLeaderFromSuccessor() (string, error) {
	return m.actualLeader, nil
}

func (m *MockTransporter) MsgNotification() <-chan *interfaces.ElectionMsg {
	return make(chan *interfaces.ElectionMsg)
}
