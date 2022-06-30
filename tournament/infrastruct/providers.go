package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/sirupsen/logrus"
)

func NewLogger() domain.Logger {
	return &Logrus{Logger: logrus.New()}
}

func NewDataInteract() chord.DataInteract {
	return NewNamedDataInteract("bunt")
}

func NewNamedDataInteract(name string) chord.DataInteract {
	return NewBuntDb(name)
}

func NewRingApi() chord.RingApi {
	return &RpcRing{}
}
