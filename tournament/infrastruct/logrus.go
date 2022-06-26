package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/sirupsen/logrus"
)

type Logrus struct {
}

func (log *Logrus) Errorf(format string, args ...any) {
	logrus.Errorf(format, args...)
}

func (log *Logrus) Infof(format string, args ...any) {
	logrus.Infof(format, args...)
}

func (log *Logrus) Info(msg string, args domain.LogArgs) {
	logrus.WithFields(logrus.Fields(args)).Info(msg)
}

func (log *Logrus) Error(msg string, args domain.LogArgs) {
	logrus.WithFields(logrus.Fields(args)).Error(msg)
}

func NewLogger() domain.Logger {
	return &Logrus{}
}
