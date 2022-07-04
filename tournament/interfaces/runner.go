package interfaces

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

// Responsible for Running a Tournament
type Runner interface {
	// Run the tournament
	Run(tm *Runnable)
}

type Runnable interface {
	GetMatches() <-chan *domain.MatchToRun
	AlreadyRun(match *domain.Pairing) (bool, domain.MatchResult)
}

type IMatchRunner interface {
	RunMatch(match *domain.Pairing) (domain.MatchResult, error)
}
