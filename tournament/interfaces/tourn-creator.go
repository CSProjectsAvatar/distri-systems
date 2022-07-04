package interfaces

import (
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"time"
)

type TournFile struct {
	Name   string
	Data   []byte
	IsGame bool
}

func SaveTournament(
	tourName string,
	tourType domain.TourType,
	files []*TournFile,
	dataMngr usecases.DataMngr,
) (*domain.TournInfo, error) {

	fileContents := make(map[string]string)
	gameFile := ""
	var players []*domain.Player
	for _, file := range files {
		fileContents[file.Name] = string(file.Data)

		if file.IsGame {
			gameFile = file.Name
		} else {
			players = append(players, &domain.Player{Id: file.Name})
		}
	}
	tourId := fmt.Sprintf("%s_%d_%v", tourName, tourType, time.Now().Unix())
	if err := dataMngr.SaveFiles(tourId, &fileContents); err != nil {
		return nil, err
	}
	info := &domain.TournInfo{
		ID:      tourId,
		Name:    gameFile,
		Players: players,
		Type_:   tourType,
	}
	return info, dataMngr.SetTournInfo(info)
}
