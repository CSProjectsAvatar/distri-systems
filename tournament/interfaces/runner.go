package interfaces

// Responsible for Running a Tournament
type runner interface {
	Run(tour_name string) error
	GetPlayGraph(tour_name string) // @todo what is the return type? Graph?

}
