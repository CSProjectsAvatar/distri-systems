package main

import (
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
)

func main() {
	myCh := make(chan int)

	fmt.Println(cap(myCh))
	fmt.Println(len(myCh))

	dm := &infrastruct.CentDataManager{}
	tMngr := usecases.NewRndTour(dm)

	runner := &infrastruct.CentRunner{}

	runner.Run(tMngr)

	myCh <- 1
}
