package tests

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	logrus "github.com/sirupsen/logrus"
	"log"
	"testing"
)

func TestMainRoutInit(t *testing.T) {
	logrus.Println("> TestMainRoutInit <")
	//localip := "127.0.0.1"
	//remote := chord.RemoteNode{
	//	Ip: "",
	//	Port: 50051,
	//}
	mr := infrastruct.NewMainRoutine(nil)
	log.Print(mr)
}
