package mocking

import (
	"log"
	"math/rand"
	"sync"
	"time"

	do "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
)

// Mock WorkerMngr
type MockWorkerMngr struct {
	RunedMatchs         map[string]int
	DrownedMatchs       map[string]int
	notificationChannel chan *do.Pairing
	initTime            time.Time

	DrownedProb float64
	lock        *sync.RWMutex
}

func NewMockWorkerMngr() *MockWorkerMngr {
	return &MockWorkerMngr{
		RunedMatchs:         make(map[string]int),
		DrownedMatchs:       make(map[string]int),
		notificationChannel: make(chan *do.Pairing),
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
			wm.DrownedMatchs[match.GetId()]++
			log.Println("Worker Drowned match", match.Player1, "and", match.Player2, "at", time.Now().Sub(wm.initTime).Seconds())
		} else {

			if coin == 0 {
				match.Winner = do.Player1Wins
			} else {
				match.Winner = do.Player2Wins
			}
			wm.notificationChannel <- match
		}
		wm.RunedMatchs[match.GetId()] += 1

	}()
}

func (wm *MockWorkerMngr) NotificationChannel() chan *do.Pairing {
	return wm.notificationChannel
}
