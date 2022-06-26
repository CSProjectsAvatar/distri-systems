package infrastruct

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/sirupsen/logrus"
	"os"
)

type Logrus struct {
	*logrus.Logger
}

func (log *Logrus) Errorf(format string, args ...any) {
	log.Logger.Errorf(format, args...)
}

func (log *Logrus) Infof(format string, args ...any) {
	log.Logger.Infof(format, args...)
}

func (log *Logrus) Info(msg string, args domain.LogArgs) {
	log.WithFields(logrus.Fields(args)).Info(msg)
}

func (log *Logrus) Error(msg string, args domain.LogArgs) {
	log.WithFields(logrus.Fields(args)).Error(msg)
}

func (log *Logrus) ToFile() domain.Logger {
	file, err := os.OpenFile(
		"logrus.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		panic(err)
	}
	log.Out = file
	log.Info("New Logger Session.", nil)
	return log
}

func NewLogger() domain.Logger {
	return &Logrus{Logger: logrus.New()}
}
