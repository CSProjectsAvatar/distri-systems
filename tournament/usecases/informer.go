package usecases

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

type Stats struct {
	Matches uint

	// Victories per player.
	Victories map[string]uint

	// Player with more victories.
	BestPlayer string
}

func GetStats(tournId string, dataMngr DataMngr) (*Stats, error) {
	matches, err := dataMngr.Matches(tournId)
	if err != nil {
		return nil, err
	}

	stats := &Stats{
		Matches:   uint(len(matches)),
		Victories: make(map[string]uint),
	}

	var bestVictories uint
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
