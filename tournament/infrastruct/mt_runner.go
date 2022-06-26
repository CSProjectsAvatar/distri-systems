package infrastruct

import (
	"fmt"
	"time"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
)

type MTRunner struct {
	WorkerMngr interfaces.IWorkerMngr
	DataMngr   interfaces.DataMngr // @todo for not repeat the games

	matchPool      chan *domain.MatchToRun
	matchWaitList  map[uint64]*domain.MatchToRun // the key its the hash of the match
	pairingResults chan *domain.Pairing
}

func NewMTRunner(workerMngr interfaces.IWorkerMngr, dataMngr interfaces.DataMngr) *MTRunner {
	runner := &MTRunner{
		WorkerMngr:     workerMngr,
		DataMngr:       dataMngr,
		matchPool:      make(chan *domain.MatchToRun, 10),
		matchWaitList:  make(map[uint64]*domain.MatchToRun),
		pairingResults: workerMngr.NotificationChannel(),
	}
	go runner.deliverMatchs()
	go runner.checkDrownedMatchs()
	go runner.receiveMatchResult()
	return runner
}

func (runner *MTRunner) Run(tm interfaces.Runnable) {
	for match := range tm.GetMatches() {
		// @todo log the match
		fmt.Println("Playing", match.Pairing.Player1, "and", match.Pairing.Player2)

		if run, result := tm.AlreadyRun(match.Pairing); run { // Match already run
			fmt.Println("Match already run")
			match.Result(result)
		} else { // Send the match to the worker pool
			runner.matchPool <- match
		}

		time.Sleep(time.Second * 2)
	}
	fmt.Println("Finished")
}

func (runner *MTRunner) deliverMatchs() { // WatchDog
	for match := range runner.matchPool {
		runner.WorkerMngr.DeliverMatch(match.Pairing) // send the match to the worker manager
		runner.matchWaitList[match.GetId()] = match   // add the match to the wait list
		match.Expiration = time.Now().Add(time.Second * time.Duration(3*2^match.TimesRetry))
	}
}

func (runner *MTRunner) checkDrownedMatchs() { // WatchDog
	for {
		// go for each match in the wait list
		for _, match := range runner.matchWaitList {
			if match.Expiration.Before(time.Now()) { // if the match is expired
				match.TimesRetry++
				runner.matchPool <- match // send the match to the pool
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func (runner *MTRunner) receiveMatchResult() { // WatchDog

	for match := range runner.pairingResults { // Result Comming
		matchRun, on_it := runner.matchWaitList[match.GetId()]
		if on_it { // Remove the match from the wait list if on it
			delete(runner.matchWaitList, match.GetId())

			matchRun.Result(match.Winner)    // Pass the result to the match
			runner.DataMngr.SaveMatch(match) // Save the match on DB
		}
	}
}
