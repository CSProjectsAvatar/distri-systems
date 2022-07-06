package main

import (
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/utils"
	"os"
	"os/signal"
)

func main() {
	var entry *chord.RemoteNode
	log := infrastruct.NewLogger().WithLevel(domain.Debug)
	if len(os.Args) > 1 { // there are program args: worker.go ip port. Program name is included.
		var port uint
		_, err := fmt.Sscan(os.Args[2], &port)
		if err != nil {
			panic(err)
		}
		entry = &chord.RemoteNode{Ip: os.Args[1], Port: port}
		log.Debug("entry set", domain.LogArgs{"entry": *entry})
	}
	infrastruct.NewMainRoutine(entry)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	log.Info(
		"worker up; stop it by pressing Ctrl+C",
		domain.LogArgs{
			"IP": utils.GetIPString(),
		})
	<-c
}
