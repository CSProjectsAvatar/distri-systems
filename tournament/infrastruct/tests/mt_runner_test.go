package infrastruct

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"

	do "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	inf "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	use "github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/stretchr/testify/assert"
)

func TestMTRunnerSimpleRun(t *testing.T) {
	log.Println("> TestMTRunner_SimpleRun <")
	workerMngr, tourMngr, runner := Init()
	// workerMngr.DrownedProb = 1

	runner.Run(tourMngr)

	expected := 4 // matches count of the MockTree
	ans := len(workerMngr.runedMatchs)

	assert.Equal(t, expected, ans, "Expected %d matches, got %d", expected, ans)
}

func TestDrownedAndRetryFunction(t *testing.T) {
	assert := assert.New(t)
	log.Println("> TestDrownedFunction <")
	workerMngr, tourMngr, runner := Init()
	runner.Run(tourMngr)

	// check that every match call the workerMngr one time more than the times it drawned
	for id, drowned := range workerMngr.drownedMatchs {
		assert.Equal(workerMngr.runedMatchs[id], drowned+1, "Expected %d times, got %d", drowned+1, workerMngr.runedMatchs[id])
	}
	assert.NotNil(tourMngr.Winner, "Expected a winner, got nil")
}

func TestThreeTournamentsAtTheSameTime(t *testing.T) {
	log.Println("> TestThreeTournamentsAtTheSameTime <")
	wm, tourMngr, runner := Init()
	wm.DrownedProb = 0.1
	tourMngr2 := use.NewRndTour(&inf.CentDataManager{})
	tourMngr3 := use.NewRndTour(&inf.CentDataManager{})

	//wait groug
	wg := &sync.WaitGroup{}
	wg.Add(3)
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
	if tourMngr.Winner == nil || tourMngr2.Winner == nil || tourMngr3.Winner == nil {
		t.Errorf("Expected a winner, got nil")
	}
}

func Init() (*MockWorkerMngr, *use.TournMngr, *inf.MTRunner) {
	wm := NewMockWorkerMngr()
	dm := &inf.CentDataManager{}
	tMngr := use.NewRndTour(dm)
	runner := inf.NewMTRunner(wm, dm)
	return wm, tMngr, runner
}

// Mock WorkerMngr
type MockWorkerMngr struct {
	notificationChannel chan *do.Pairing
	runedMatchs         map[uint64]int
	drownedMatchs       map[uint64]int
	initTime            time.Time

	DrownedProb float64
	lock        *sync.RWMutex
}

func NewMockWorkerMngr() *MockWorkerMngr {
	return &MockWorkerMngr{
		notificationChannel: make(chan *do.Pairing),
		runedMatchs:         make(map[uint64]int),
		drownedMatchs:       make(map[uint64]int),
		lock:                &sync.RWMutex{},
		DrownedProb:         0.2,
		initTime:            time.Now(),
	}
}
func (wm *MockWorkerMngr) DeliverMatch(match *do.Pairing) {
	go func() {
		// time.Sleep(time.Second * 1)
		// random
		coin := time.Now().UnixMilli() % 2
		// random uniform distribution
		coin2 := rand.Float64()
		wm.lock.Lock()
		defer wm.lock.Unlock()

		if coin2 < wm.DrownedProb {
			wm.drownedMatchs[match.GetId()]++
			log.Println("Worker Drowned match", match.Player1, "and", match.Player2, "at", time.Now().Sub(wm.initTime).Seconds())
		} else {

			if coin == 0 {
				match.Winner = do.Player1Wins
			} else {
				match.Winner = do.Player2Wins
			}
			wm.notificationChannel <- match
		}
		wm.runedMatchs[match.GetId()] += 1

	}()
}

func (wm *MockWorkerMngr) NotificationChannel() chan *do.Pairing {
	return wm.notificationChannel
}
