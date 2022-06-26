package domain

import (
	"fmt"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/utils"
)

type Pairing struct {
	ID uint64

	TourId  int
	Player1 *Player
	Player2 *Player

	Winner MatchResult
}

// @todo NewPairing builder

func (m *Pairing) GetId() uint64 { // @audit-issue not having in count 2 match with the same pairing
	// Calculate a hash
	if m.ID == 0 {
		m.ID = utils.Hash(fmt.Sprintf("%d%d%d", m.TourId, m.Player1.Name, m.Player2.Name))
	}
	return m.ID
}

type MatchToRun struct {
	Pairing *Pairing
	result  chan MatchResult

	Expiration time.Time
	TimesRetry int
}

func (m *MatchToRun) Result(res MatchResult) {
	m.Pairing.Winner = res
	m.result <- res
}

func (m *MatchToRun) GetId() uint64 {
	return m.Pairing.GetId()
}

type Player struct {
	Name string
}

type TournInfo struct {
	Name          string
	Type_         TourType
	Players       []Player
	AvrgMatchTime time.Duration
}
