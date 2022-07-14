package infrastruct

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

type MTRunner struct {
	WorkerMngr interfaces.IWorkerMngr
	DataMngr   usecases.DataMngr // @todo for not repeat the games

	matchPool      chan *domain.MatchToRun
	pairingResults <-chan *domain.Pairing
	matchWaitList  map[string]*domain.MatchToRun // the key its the hash of the match

	waitList *sync.RWMutex
}

func NewMTRunner(workerMngr interfaces.IWorkerMngr, dataMngr usecases.DataMngr) *MTRunner {
	runner := &MTRunner{
		WorkerMngr:     workerMngr,
		DataMngr:       dataMngr,
		matchPool:      make(chan *domain.MatchToRun, 10),
		matchWaitList:  make(map[string]*domain.MatchToRun),
		pairingResults: workerMngr.NotificationChannel(),
		waitList:       &sync.RWMutex{},
	}
	go runner.deliverMatchs()
	go runner.checkDrownedMatchs()
	go runner.receiveMatchResult()
	return runner
}

func (runner *MTRunner) Run(tm interfaces.Runnable) {
	for match := range tm.GetMatches() {
		if run, result := tm.AlreadyRun(match.Pairing); run { // Match already run
			log.Println("Match already run")
			match.Result(result)
		} else { // Send the match to the worker pool
			runner.matchPool <- match
		}
	}
	log.Println("Finished Tour")
}

func (r *MTRunner) deliverMatchs() { // WatchDog
	for match := range r.matchPool {
		log.Println("Trying to get Playing", match.Pairing.Player1, "and", match.Pairing.Player2)
		r.WorkerMngr.DeliverMatch(match.Pairing) // send the match to the worker manager

		r.waitList.Lock()                      // [+]
		r.matchWaitList[match.GetId()] = match // add the match to the wait list
		addedT := time.Duration(math.Pow(2, float64(match.TimesRetry)))
		addedT++
		match.Expiration = time.Now().Add(domain.BaseWhaitTime * addedT)
		r.waitList.Unlock() // [-]
	}
}

func (r *MTRunner) checkDrownedMatchs() { // WatchDog
	for {
		// go for each match in the wait list
		r.waitList.RLock() // [+]
		for _, match := range r.matchWaitList {
			if match.Expiration.Before(time.Now()) { // if the match is expired
				match.TimesRetry++
				r.matchPool <- match // send the match to the pool
				log.Println("Match", match.Pairing.Player1, "and", match.Pairing.Player2, "drowned")
			}
		}
		r.waitList.RUnlock() // [-]
		time.Sleep(domain.BaseWhaitTime)
	}
}

func (r *MTRunner) receiveMatchResult() { // WatchDog

	for match := range r.pairingResults { // Result Comming
		r.waitList.RLock() // [+]
		matchRun, on_it := r.matchWaitList[match.GetId()]
		r.waitList.RUnlock() // [-]
		if on_it {           // Remove the match from the wait list if on it
			r.waitList.Lock() // [+]
			delete(r.matchWaitList, match.GetId())
			r.waitList.Unlock() // [-]

			log.Println("Match", matchRun.Pairing.Player1, "and", matchRun.Pairing.Player2, "result:", match.Winner)
			matchRun.Result(match.Winner)  // Pass the result to the match
			go r.DataMngr.SaveMatch(match) // Save the match on DB
		}
	}
}
