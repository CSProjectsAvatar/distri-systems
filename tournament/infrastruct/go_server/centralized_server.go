// // Centralized go Server that implements UploadTournament
package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	// "github.com/CSProjectsAvatar/distri-systems/go_server/pb"
// )

// type server struct {
// 	pb.UnimplementedMiddlewareServer
// }

// func (*server) UploadTournament(ctx context.Context, req *pb.TournamentReq) (*pb.TournamentResp, error) {
// 	fmt.Println("RPC UploadTournament Called on Server")
// 	// save the tournament files on the folder "files/tournament_name"
// 	t_name := req.Name

// 	return nil, nil
// }

// func (*server) RunTournament(req *pb.RunReq, stream *pb.Middleware_RunTournamentServer) error {
// 	fmt.Println("RPC RunTournament Called on Server")
// 	return nil
// }

// func main() {
//     b, err := os.ReadFile("input.txt")
//     if err != nil {
//         log.Fatal(err)
//     }

//     // `b` contains everything your file has.
//     // This writes it to the Standard Out.
//     os.Stdout.Write(b)

//     // You can also write it to a file as a whole.
//     err = os.WriteFile("destination.txt", b, 0644)
//     if err != nil {
//         log.Fatal(err)
//     }
// }

// func main() {

// 	out, err := RunMatch("coin", "p1.py", "p2.py")
// 	if err != nil {
// 		log.Printf("error: %v\n", err)
// 	}
// 	fmt.Println("--- stdout ---")
// 	fmt.Println(out)
// }
