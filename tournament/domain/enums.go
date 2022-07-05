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
	BaseWhaitTime         = time.Second * 3
	WhaitTimeBetweenRetry = time.Second * 2
	MaxRetryTimes         = 3
)

const IdLength = 56
const ReplicaNumber = 4

const ChordPort = 9090
const WMngrPort = 8080
const WClientPort = 8081
const MiddPort = 8082
