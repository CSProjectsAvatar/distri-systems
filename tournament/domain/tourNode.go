package domain

import (
	"fmt"
	"sort"
	"sync"

	"github.com/CSProjectsAvatar/distri-systems/utils"
)

// TourNode is the Tournament Node
type TourNode struct {
	Children []*TourNode // If I am a leaf, I am a player
	Winner   *Player

	JoinFunc func(childWinners <-chan *Player, winnerCh chan<- *Player) []*MatchToRun
}

func DefNodeFunc(childWinners <-chan *Player, winnerCh chan<- *Player) []*MatchToRun {
	winnerSlice := make([]*Player, 0)
	for len(winnerSlice) < cap(childWinners) {
		winnerSlice = append(winnerSlice, <-childWinners)
	} // Whait for all the winners, @audit-info can't return more than one winner per group

	wg := &sync.WaitGroup{}
	gamesWon := make([]float64, len(winnerSlice))
	matches := make([]*MatchToRun, 0)
	for i, pi := range winnerSlice {
		// get the players ahead
		for j := i + 1; j < len(winnerSlice); j++ {
			pj := winnerSlice[j]

			matchToRun := NewMatchToRun(pi, pj)
			matches = append(matches, matchToRun) // Add the match to the list of matches to run
			wg.Add(1)                             // Add the match to the wait group

			// Keeps track of the games won by each player
			go func(i, j int) {
				defer wg.Done()

				res := <-matchToRun.result
				if res == Player1Wins {
					gamesWon[i]++
				} else if res == Player2Wins {
					gamesWon[j]++
				} else if res == Draw {
					gamesWon[i] += 0.5
					gamesWon[j] += 0.5
				}
			}(i, j)
		}
	}
	// Send the best Winner to the winner channel
	go func() {
		defer close(winnerCh)
		wg.Wait()
		PassBestWinner(gamesWon, winnerSlice, winnerCh)
	}()
	return matches
}

func PassBestWinner(gamesWon []float64, winnerSlice []*Player, winnerCh chan<- *Player) {
	tuples := utils.JoinToTuples(gamesWon, winnerSlice)
	// sort the tuples
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].First > tuples[j].First
	})
	// send the best player to the winner channel
	winnerCh <- tuples[0].Second
}

func PassBestHalfWinners(gamesWon []float64, winnerSlice []*Player, winnerCh chan<- *Player) {
	tuples := utils.JoinToTuples(gamesWon, winnerSlice)
	// sort the tuples
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].First > tuples[j].First
	})
	// send the best half of players to the winner channel
	for i := 0; i < len(tuples)/2; i++ {
		winnerCh <- tuples[i].Second
	}
}

// runnerCh: Channel for the runner run the matchs
// winnerCh: Channel for the father to get the Winner
func (tnode *TourNode) PlayNode(runnerCh chan<- *MatchToRun, winnerCh chan<- *Player) {
	if len(tnode.Children) == 0 {
		winnerCh <- tnode.Winner
		fmt.Println("Leaf ", tnode.Winner) //@todo remove
		return
	}
	// Processing Children
	childWinners := make(chan *Player, len(tnode.Children))
	for _, child := range tnode.Children {
		go child.PlayNode(runnerCh, childWinners)
	}
	matches := tnode.JoinFunc(childWinners, winnerCh)
	fmt.Println("Matches to run: ", len(matches)) //@todo remove
	// Send the matches to the runner
	for _, match := range matches {
		runnerCh <- match
	}
	//close(childWinners)
}
