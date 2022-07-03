package domain

import "time"

type TourType int

const (
	First_Defeat TourType = iota
	All_vs_All
)

type MatchResult int

const (
	NotPlayed MatchResult = iota
	Player1Wins
	Player2Wins
	Draw
)

var (
	BaseWhaitTime = time.Second * 3
)
