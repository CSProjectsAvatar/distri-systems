package infrastruct

import (
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	"time"
)

type CentRunner struct {
}

func (runner *CentRunner) Run(tm interfaces.Runnable) {
	matches := tm.GetMatches()

	for {
		match := <-matches
		// Players
		fmt.Println("Playing", match.Pairing.Player1, "and", match.Pairing.Player2)

		// Mock the match
		// @todo call dm if already ran, only one time at the beguining
		match.Result <- domain.Player2Wins
		// @todo call dm for save
		fmt.Println("Winner: ", match.Pairing.Player2)

		time.Sleep(time.Second * 2)
	}
	fmt.Println("Finished")
}
