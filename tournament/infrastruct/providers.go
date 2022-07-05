package infrastruct

import (
	"crypto/sha1"
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/sirupsen/logrus"
	"time"
)

func NewLogger() domain.Logger {
	return &Logrus{Logger: logrus.New()}
}

func NewDataInteract() chord.DataInteract {
	return NewNamedDataInteract("bunt")
}

func NewTestDht[T any](replicas uint, remote *chord.RemoteNode) *usecases.Dht[T] {
	return usecases.NewDhtBuilder[T]().
		Ring(NewRingApi(remote)).
		Log(NewLogger().ToFile()).
		Replicas(replicas).
		Remote(remote).
		Build()
}

func NewNamedDataInteract(name string) chord.DataInteract {
	return NewBuntDb(name)
}

func NewRingApi(client *chord.RemoteNode) chord.RingApi {
	return &RpcRing{
		client: client,
		log:    NewLogger(),
	}
}

func ChordConfig(ip string, port uint) *chord.Config {
	return &chord.Config{
		Ip:   ip,
		Port: port,
		Hash: sha1.New,
		Ring: NewRingApi(&chord.RemoteNode{Ip: ip, Port: port}),
		Data: NewNamedDataInteract(
			fmt.Sprintf("bunt-%d-%v", port, time.Now().UnixMilli())),
		M:           domain.IdLength,
		IncludeDate: true,
	}
}
func LocalChordConfig(port uint) *chord.Config {
	return ChordConfig("127.0.0.1", port)
}
