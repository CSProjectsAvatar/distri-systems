package tests

import (
	"testing"

	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/stretchr/testify/assert"
)

func TestTwoMatchWithSamePairing(t *testing.T) {
	assert := assert.New(t)
	// dm := &mocking.CentDataManager{}
	tm := usecases.NewMockTourMngr()

	p1 := tm.TInfo.Players[0]
	p2 := tm.TInfo.Players[1]
	p3 := tm.TInfo.Players[2]
	p4 := tm.TInfo.Players[3]

	m1 := tm.GetMatch(p1, p2)
	m2 := tm.GetMatch(p3, p4)
	assert.NotEqual(m1.GetId(), m2.GetId())

	m1_2 := tm.GetMatch(p1, p2)
	m2_2 := tm.GetMatch(p3, p4)
	assert.NotEqual(m1.GetId(), m1_2.GetId())
	assert.NotEqual(m2.GetId(), m2_2.GetId())
}
