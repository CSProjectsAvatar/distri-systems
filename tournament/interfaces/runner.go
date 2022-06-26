package interfaces

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

// Responsible for Running a Tournament
type Runner interface {
	// Run the tournament
	Run(tm *Runnable)

	// AlreadyRan Returns True if a match has already been run
	AlreadyRan(match *domain.Pairing) bool

	// MatchRan Marks a match as ran
	MatchRan(match *domain.Pairing)
}

type Runnable interface {
	GetMatches() <-chan *domain.MatchToRun
	AlreadyRun(match *domain.Pairing) (bool, domain.MatchResult)
}
