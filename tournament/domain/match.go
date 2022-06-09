package domain

type Match struct {
	Player1 *Player
	Player2 *Player
	Winner  MatchResult
}

type MatchToRun struct {
	Pairing Match
	Result  chan MatchResult
}

type Player struct {
	Name string
}

type TournInfo struct {
	Name    string
	Type_   TourType
	Players []Player
}
