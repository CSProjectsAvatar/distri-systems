package interfaces

import "github.com/CSProjectsAvatar/distri-systems/tournament/domain"

// Responsible for Running a Tournament
type Runner interface {
	Run(tm *Runnable)
}

type Runnable interface {
	GetMatches() <-chan *domain.MatchToRun
}
