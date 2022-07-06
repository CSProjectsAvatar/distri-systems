package tests

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	inf "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	mock "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/mocking"
	use "github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/stretchr/testify/assert"
)

func TestMTRunnerSimpleRun(t *testing.T) {
	log.Println("> TestMTRunner_SimpleRun <")
	workerMngr, tourMngr, runner, _ := Init()
	// workerMngr.DrownedProb = 1

	runner.Run(tourMngr)

	expected := 4 // matches count of the MockTree
	ans := len(workerMngr.RunedMatchs)

	assert.Equal(t, expected, ans, "Expected %d matches, got %d", expected, ans)
}

func TestDrownedAndRetryFunction(t *testing.T) {
	assert := assert.New(t)
	log.Println("> TestDrownedFunction <")
	workerMngr, tourMngr, runner, _ := Init()
	runner.Run(tourMngr)

	// check that every match call the workerMngr one time more than the times it drawned
	for id, drowned := range workerMngr.DrownedMatchs {
		assert.Equal(workerMngr.RunedMatchs[id], drowned+1, "Expected %d times, got %d", drowned+1, workerMngr.RunedMatchs[id])
	}
	assert.NotNil(tourMngr.TInfo.Winner, "Expected a winner, got nil")
}

func TestThreeTournamentsAtTheSameTime(t *testing.T) {
	log.Println("> TestThreeTournamentsAtTheSameTime <")
	wm, tourMngr, runner, dm := Init()
	wm.DrownedProb = 0.1
	tourMngr2 := use.NewMockTourMngr()
	tourMngr3 := use.NewMockTourMngr()
	tourMngr2.SetDM(dm)
	tourMngr3.SetDM(dm)
	//wait groug
	wg := &sync.WaitGroup{}
	wg.Add(3)
	//wg.Add(2)
	go func() {
		runner.Run(tourMngr)
		wg.Done()
	}()
	go func() {
		runner.Run(tourMngr2)
		wg.Done()
	}()
	go func() {
		runner.Run(tourMngr3)
		wg.Done()
	}()

	wg.Wait()
	if tourMngr.TInfo.Winner == nil || tourMngr2.TInfo.Winner == nil || tourMngr3.TInfo.Winner == nil {
		//if tourMngr.TInfo.Winner == nil || tourMngr2.TInfo.Winner == nil {
		t.Errorf("Expected a winner, got nil")
	}
}

func Init() (*mock.MockWorkerMngr, *use.TournMngr, *inf.MTRunner, use.DataMngr) {
	wm := mock.NewMockWorkerMngr()
	dm := &mock.CentDataManager{}
	tMngr := use.NewMockTourMngr()
	tMngr.SetDM(dm)
	runner := inf.NewMTRunner(wm, dm)
	domain.BaseWhaitTime = 1 * time.Second
	return wm, tMngr, runner, dm
}
