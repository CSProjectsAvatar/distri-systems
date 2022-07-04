package transport

import (
	"context"

	. "github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	pb_r "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_ring"
	pb_w "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_workerMngr"
	. "github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	log "github.com/sirupsen/logrus"
)

type WorkerTransport struct {
	BaseTransport
	pb_r.UnimplementedRingServer
	leadProv ILeaderProvider
	sucProv  ISuccProvider

	msgsChan chan *ElectionMsg
}

func NewWorkerClient(config *Config, leadProv ILeaderProvider, sucProv ISuccProvider) (*WorkerTransport, error) {
	t, err := NewBaseTransport(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	wT := &WorkerTransport{
		BaseTransport: *t,
		leadProv:      leadProv,
		sucProv:       sucProv,
		msgsChan:      make(chan *ElectionMsg, 1),
	}

	// Register the server
	pb_r.RegisterRingServer(wT.server, wT)

	return wT, nil
}

func (wT *WorkerTransport) getConnWM(addr string) (pb_w.WorkerMngrClient, error) {
	remote, err := wT.BaseTransport.getConn(addr)
	if err != nil {
		return nil, err
	}
	client := pb_w.NewWorkerMngrClient(remote.conn)
	return client, nil
}

func (wT *WorkerTransport) SendResults(match *Pairing) error {
	addr := wT.leadProv.GetLeader()
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
	addr := wT.leadProv.GetLeader()
	client, err := wT.getConnWM(addr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), wT.config.Timeout)
	defer cancel()

	resp, err := client.GiveMeWork(ctx, &pb_w.MatchReq{})
	if err != nil {
		return nil, err
	}
	return FromMatchResp(resp), nil
}

//							 ^
// WorkerMngr Client Methods |

func (wT *WorkerTransport) getConnRng(addr string) (pb_r.RingClient, error) {
	remote, err := wT.BaseTransport.getConn(addr)
	if err != nil {
		return nil, err
	}
	client := pb_r.NewRingClient(remote.conn)
	return client, nil
}

// Client
func (wT *WorkerTransport) SendToSuccessor(msg *ElectionMsg) error {
	addr := wT.sucProv.GetSuccessor()
	client, err := wT.getConnRng(addr)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), wT.config.Timeout)
	defer cancel()

	req := NewElectionMsgReq(msg)
	_, err = client.SendMessage(ctx, req)

	return err
}

func (wT *WorkerTransport) GetLeaderFromSuccessor() (string, error) {
	addr := wT.sucProv.GetSuccessor()
	client, err := wT.getConnRng(addr)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), wT.config.Timeout)
	defer cancel()

	resp, err := client.GetLeader(ctx, &pb_r.GetLeaderReq{})
	if err != nil {
		return "", err
	}
	return resp.LeaderAddr, nil
}

// Server
func (wT *WorkerTransport) MsgNotification() <-chan *ElectionMsg {
	return wT.msgsChan
}

// Message Received on Server
func (wT *WorkerTransport) SendMessage(ctx context.Context, in *pb_r.ElectionMsgReq) (*pb_r.ElectionMsgResp, error) {
	wT.msgsChan <- &ElectionMsg{
		Type:      ElectionType(in.Type),
		OnTheRing: in.OnIt,
	}
	return &pb_r.ElectionMsgResp{}, nil
}

func (wT *WorkerTransport) GetLeader(ctx context.Context, in *pb_r.GetLeaderReq) (*pb_r.GetLeaderResp, error) {
	addr := wT.leadProv.GetLeader()
	return &pb_r.GetLeaderResp{LeaderAddr: addr}, nil
}

// 						  ^
// IRingTransport methods |
