package transport

import (
	"context"
	"math/rand"
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
	log.Println("Called UploadTournament, tourName: ", in.Name)
	tInfo, err := interfaces.SaveTournament(in.Name, domain.TourType(in.TourType), ToTournFiles(in.Files), mid.dm)
	if err != nil {
		return nil, err
	}
	return &pb_m.TournamentResp{
		TourId: tInfo.ID,
	}, nil
}
func (mid *MidServer) GetStats(ctx context.Context, in *pb_m.StatsReq) (*pb_m.StatsResp, error) {
	log.Println("Called GetStats")
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
	log.Println("Called GetAllIds")
	ids, err := mid.dm.GetAllIds()
	if err != nil {
		return nil, err
	}
	return &pb_m.AllIdsResp{
		TourIds: ids,
	}, nil
}

func (mid *MidServer) GetRndStats(ctx context.Context, in *pb_m.StatsReq) (*pb_m.StatsResp, error) {
	log.Println("Called GetRndStats")
	rndStat := &pb_m.StatsResp{
		// rand betwen 0 and 15
		Matches:    uint32(rand.Intn(15)),
		BestPlayer: "Player2",
		TourName:   "Chez",
	}
	return rndStat, nil
}
