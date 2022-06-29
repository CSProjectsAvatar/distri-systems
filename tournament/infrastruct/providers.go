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
	return usecases.NewDataMap()
}

func NewRingApi() chord.RingApi {
	return &RpcRing{}
}
