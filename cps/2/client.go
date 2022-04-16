package main

import (
	"log"
	"net/rpc"
	"os"

	"github.com/CSProjectsAvatar/distri-systems/cps/2/calc_rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	callDoc(client)
	callIam(client, os.Args[1])
	callMul(client)
}

func callDoc(client *rpc.Client) {
	var reply string
	var args calc_rpc.Args
	err := client.Call("RemoteCalc.Doc", &args, &reply)
	if err != nil {
		log.Fatal("RemoteCalc error:", err)
	}
	log.Printf("RemoteCalc.Doc: \n%v", reply)
}

func callIam(client *rpc.Client, name string) {
	var reply string
	err := client.Call("RemoteCalc.Iam", &name, &reply)
	if err != nil {
		log.Fatal("RemoteCalc error:", err)
	}
	log.Printf("RemoteCalc: %v", reply)
}

func callMul(client *rpc.Client) {
	args := &calc_rpc.Args{7, 8}
	var reply float64
	err := client.Call("RemoteCalc.Mul", args, &reply)
	if err != nil {
		log.Fatal("RemoteCalc error:", err)
	}
	log.Printf("RemoteCalc: %v*%v=%v\n", args.A, args.B, reply)
}
