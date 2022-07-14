package transport

import (
	"context"
	"regexp"
	"strconv"

	//"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"

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
	//config.Port = WClientPort
	t, err := NewBaseTransport(config)
	if err != nil {
		log.Fatal("Couldn't create base transport", err)
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
	addr := AssureSrvAddress(wT.leadProv.GetLeader())
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
	addr := AssureSrvAddress(wT.leadProv.GetLeader())
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
	addr := AssureCltAddress(wT.sucProv.GetSuccessor())
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
	addr := AssureCltAddress(wT.sucProv.GetSuccessor())
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

func AssureSrvAddress(address string) string {
	if itsFullAddress(address) {
		return address
	}
	return address + ":" + strconv.Itoa(WMngrPort)
}

func AssureCltAddress(address string) string {
	if itsFullAddress(address) {
		return address
	}
	return address + ":" + strconv.Itoa(WClientPort)
}

// if its the full address like *:*
func itsFullAddress(address string) bool {
	// if address has : return it, else atach it the port
	regex := `^(.*):(\d+)$`
	if matched, _ := regexp.MatchString(regex, address); matched {
		return true
	}
	return false
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
