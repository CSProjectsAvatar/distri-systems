package usecases

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

type Stats struct {
	Matches uint32

	// Victories per player.
	Victories map[string]uint32

	// Player with more victories.
	BestPlayer string

	// Winner of the tournament.
	Winner string

	Name string
}

func EmptyStat() *Stats {
	return &Stats{
		Victories:  make(map[string]uint32),
		Matches:    0,
		BestPlayer: "",
		Winner:     "",
		Name:       "",
	}
}
func GetStats(tournId string, dataMngr DataMngr) (*Stats, error) {
	matches, err := dataMngr.Matches(tournId)
	if err != nil {
		return EmptyStat(), nil
	}
	tInfo, err := dataMngr.GetTournInfo(tournId)
	if err != nil {
		return EmptyStat(), nil
	}

	winner := ""
	if tInfo.Winner != nil {
		winner = tInfo.Winner.Id
	}
	stats := &Stats{
		Matches:   uint32(len(matches)),
		Victories: make(map[string]uint32),
		Winner:    winner,
		Name:      tInfo.Name,
	}

	var bestVictories uint32
	for _, match := range matches {
		player := ""
		if match.Winner == domain.Player1Wins {
			player = match.Player1.Id
		} else if match.Winner == domain.Player2Wins {
			player = match.Player2.Id
		}
		if player != "" {
			victs := stats.Victories[player]
			stats.Victories[player] = victs + 1
			victs++

			if victs > bestVictories {
				bestVictories = victs
				stats.BestPlayer = player
			}
		}
	}

	return stats, nil
}
