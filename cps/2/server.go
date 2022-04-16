package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/CSProjectsAvatar/distri-systems/cps/2/calc_rpc"
)

func main() {
	calc := new(calc_rpc.RemoteCalc)
	rpc.Register(calc)

	rpc.HandleHTTP()
	log.Println("Listening on port 1234 (TCP)...")
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("Serving HTTP...")
	http.Serve(l, nil)
}
