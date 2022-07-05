package transport

import (
	"context"
	"net"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"

	pb_m "github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct/pb_mid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Server Side Worker Manager
type MidServer struct {
	dm     usecases.DataMngr
	addr   string
	sock   *net.TCPListener
	server *grpc.Server
	pb_m.UnimplementedMiddlewareServer
}

func NewMidServer(addr string, dataM usecases.DataMngr) *MidServer {
	mid := &MidServer{
		dm:   dataM,
		addr: addr,
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	mid.sock = lis.(*net.TCPListener)
	mid.server = grpc.NewServer()

	pb_m.RegisterMiddlewareServer(mid.server, mid)
	return mid
}

func (mid *MidServer) Start() error {
	go mid.server.Serve(mid.sock)
	return nil
}

func (mid *MidServer) Stop() error {
	mid.server.Stop()
	mid.sock.Close()
	return nil
}

// Saves the tournament files
func (mid *MidServer) UploadTournament(ctx context.Context, in *pb_m.TournamentReq) (*pb_m.TournamentResp, error) {
	tInfo, err := interfaces.SaveTournament(in.Name, domain.TourType(in.TourType), ToTournFiles(in.Files), mid.dm)
	if err != nil {
		return nil, err
	}
	return &pb_m.TournamentResp{
		TourId: tInfo.ID,
	}, nil
}
func (mid *MidServer) GetStats(ctx context.Context, in *pb_m.StatsReq) (*pb_m.StatsResp, error) {
	stats, err := usecases.GetStats(in.TourId, mid.dm)
	if err != nil {
		return nil, err
	}
	return &pb_m.StatsResp{
		Matches:    uint32(stats.Matches),
		Victories:  stats.Victories,
		BestPlayer: stats.BestPlayer,
		Winner:     stats.Winner,
		TourName:   stats.Name,
	}, nil
}

//get all ids of tournaments
func (mid *MidServer) GetAllIds(ctx context.Context, in *pb_m.AllIdsReq) (*pb_m.AllIdsResp, error) {
	ids, err := mid.dm.GetAllIds()
	if err != nil {
		return nil, err
	}
	return &pb_m.AllIdsResp{
		TourIds: ids,
	}, nil
}

//GetStats(ctx context.Context, in *StatsReq, opts ...grpc.CallOption) (*StatsResp, error)
