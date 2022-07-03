package transport

import (
	"context"

	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	pb "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_workerMngr"
	. "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	log "github.com/sirupsen/logrus"
)

type WorkerTransport struct {
	BaseTransport
	prov ILeaderProvider
}

func NewWorkerClient(config *Config, provider ILeaderProvider) (*WorkerTransport, error) {
	t, err := NewBaseTransport(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	wT := &WorkerTransport{
		BaseTransport: *t,
		prov:          provider,
	}
	return wT, nil
}

func (wT *WorkerTransport) getConnWM(addr string) (pb.WorkerMngrClient, error) {
	remote, err := wT.BaseTransport.getConn(addr)
	if err != nil {
		return nil, err
	}
	client := pb.NewWorkerMngrClient(remote.conn)
	return client, nil
}

func (wT *WorkerTransport) SendResults(match *Pairing) error {
	addr := wT.prov.GetLeader()
	client, err := wT.getConnWM(addr)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), wT.config.Timeout)
	defer cancel()

	_, err = client.CatchResult(ctx, NewResultReq(match))
	if err != nil {
		return err
	}
	return nil
}

func (wT *WorkerTransport) GetMatchToRun() (*Pairing, error) {
	addr := wT.prov.GetLeader()
	client, err := wT.getConnWM(addr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), wT.config.Timeout)
	defer cancel()

	resp, err := client.GiveMeWork(ctx, &pb.MatchReq{})
	if err != nil {
		return nil, err
	}
	return FromMatchResp(resp), nil
}

//							 ^
// WorkerMngr Client Methods |

// // Client
// func (wT *WorkerTransport) SendToSuccessor(msg *ElectionMsg) {
// 	addr := wT.prov.GetLeader()
// 	client, err := wT.getConnR(addr)
// 	if err != nil {
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), wT.config.Timeout)
// 	defer cancel()

// 	_, err = client.SendToSuccessor(ctx, msg)
// 	if err != nil {
// 		return
// 	}
// }

// GetLeaderFromSuccessor() string
// // Server
// MsgNotification() <-chan *ElectionMsg
