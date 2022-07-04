package domain

import (
	"fmt"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/utils"
)

type Pairing struct {
	ID string

	TourId  string
	Player1 *Player
	Player2 *Player

	Winner MatchResult
}

func NewPairing(tourId string, player1, player2 *Player, pairing_time int) *Pairing {
	pair := &Pairing{
		TourId:  tourId,
		Player1: player1,
		Player2: player2,
	}
	pair.SetId(pairing_time)
	return pair
}

func (m *Pairing) SetId(pairTime int) {
	// Calculate a hash
	if m.ID == "" {
		m.ID = utils.Hash(fmt.Sprintf("%s%s%s%s", m.TourId, m.Player1.Id, m.Player2.Id, pairTime))
	}
}

func (m *Pairing) GetId() string {
	return m.ID
}

type MatchToRun struct {
	Pairing *Pairing
	result  chan MatchResult

	Expiration time.Time
	TimesRetry int
}

func NewMatchToRun(tourId string, pi, pj *Player, pairNmbr int) *MatchToRun {
	return &MatchToRun{
		Pairing: NewPairing(tourId, pi, pj, pairNmbr),
		result:  make(chan MatchResult),
	}
}

func (m *MatchToRun) Result(res MatchResult) {
	m.Pairing.Winner = res
	m.result <- res
}

func (m *MatchToRun) GetId() string {
	return m.Pairing.GetId()
}

type Player struct {
	Id string
}

type TournInfo struct {
	ID      string
	Name    string
	Type_   TourType
	Players []*Player
	Winner  *Player
}
