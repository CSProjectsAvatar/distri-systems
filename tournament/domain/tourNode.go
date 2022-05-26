package domain

type Player struct {
}

// TourNode is the Tournament Node
type TourNode struct {
	children []TourNode
	winner   Player

	processWinners func(winnerCh <-chan Player)
}

func (tnode *TourNode) Play(winnerCh chan<- Player) {
	if len(tnode.children) == 0 {
		winnerCh <- tnode.winner
		return
	}
	// Processing Children
	childWinners := make(chan Player, len(tnode.children))
	for _, child := range tnode.children {
		go child.Play(childWinners)
	}

	tnode.processWinners(childWinners)
	close(childWinners)
}
