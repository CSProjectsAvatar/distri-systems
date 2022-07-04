package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/sirupsen/logrus"
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
