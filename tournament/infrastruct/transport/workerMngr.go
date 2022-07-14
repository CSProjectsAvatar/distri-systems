package transport

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	. "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_workerMngr"
)

// Server Side Worker Manager
type WorkerMngr struct {
	UnimplementedWorkerMngrServer

	sock   *net.TCPListener
	server *grpc.Server

	matchesToRun chan *Pairing
	results      chan *Pairing
}

func NewWorkerMngr(addr string) (*WorkerMngr, error) {
	mngr := &WorkerMngr{
		matchesToRun: make(chan *Pairing),
		results:      make(chan *Pairing, 10),
	}
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Error("failed to listen: %v", err)
		return nil, err
	} else {
		log.Println("Listen on %v", addr)
	}
	mngr.sock = lis.(*net.TCPListener)
	mngr.server = grpc.NewServer()

	RegisterWorkerMngrServer(mngr.server, mngr)
	return mngr, nil
}

func (wm *WorkerMngr) Start() error {
	go wm.server.Serve(wm.sock)
	return nil
}

func (wm *WorkerMngr) Stop() error {
	wm.server.Stop()
	wm.sock.Close()
	return nil
}

func (mngr *WorkerMngr) GiveMeWork(ctx context.Context, in *MatchReq) (*MatchResp, error) {
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "The client canceled the request")
	case match := <-mngr.matchesToRun:
		log.Println("Send match to worker", match.ID)
		return &MatchResp{
			MatchId:     match.ID,
			TourId:      match.TourId,
			FstPlayerID: match.Player1.Id,
			SndPlayerID: match.Player2.Id,
		}, nil
	}
}

func (mngr *WorkerMngr) CatchResult(ctx context.Context, in *ResultReq) (*ResultResp, error) {
	log.Println("Received result", in.MatchId, "wins:", in.Winner)
	mngr.results <- &Pairing{
		ID:      in.MatchId,
		TourId:  in.TourId,
		Player1: &Player{Id: in.FstPlayerID},
		Player2: &Player{Id: in.SndPlayerID},
		Winner:  MatchResult(in.Winner),
	}
	return &ResultResp{}, nil
}
func (mngr *WorkerMngr) DeliverMatch(match *Pairing) {
	mngr.matchesToRun <- match
}
func (mngr *WorkerMngr) NotificationChannel() <-chan *Pairing {
	return mngr.results
}
