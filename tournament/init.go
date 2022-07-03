package main

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
)

func main() {

	dm := &infrastruct.CentDataManager{}
	tMngr := usecases.NewRndTour(dm)

	runner := &infrastruct.CentRunner{}

	runner.Run(tMngr)

}
