package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/CSProjectsAvatar/distri-systems/go_server/pb"
)

// struct responsible for managing the tournament
// it is responsible for creating the tournament folder and saving the files
// it is responsible for running the tournament and saving the info in "../files/tournament_name/<t_name>.info"
// the matchs are given by getMatchs(), wich accepts a tournament Type and a list of players and returns a list of matchs
type TournamentManager struct {
	t_name string
	t_type pb.TournType

	resolver Resolver
	files    map[string]string
	players  []string
}

func (t *TournamentManager) getMatchs() []pb.Match {
	var matchs []pb.Match
	switch t.t_type {
	case pb.TournType_All_vs_All:
		matchs = getAllVsAllMatchs(t.players)
	case pb.TournType_First_Defeat:
		matchs = getFirstDefeatMatchs(t.players)
	}
	return matchs
}

func getAllVsAllMatchs(players []string) []pb.Match {
	var matchs []pb.Match
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			matchs = append(matchs, pb.Match{
				Players: []string{players[i], players[j]},
			})
		}
	}
	return matchs
}

func getFirstDefeatMatchs(players []string) []pb.Match {
	// @TODO: implement
	return nil
}

// Run a match between two players by calling the game with them as parameters
func RunMatch(tourN_name, player_i, player_j string) (GameResult, error) {
	command := "python ../files/" + tourN_name + "/game.py " + player_i + " " + player_j
	err, out, _ := RunCommand(command)
	if err != nil {
		log.Printf("error: %v\n", err)
		return 0, err
	}
	winner, err := strconv.Atoi(out[0:1]) // for deal with "2\r\n" response styles
	if err != nil {
		log.Printf("Couldn parse response from %v\n error: %v", tourN_name, err)
	}
	return GameResult(winner), nil
}

// Write the tournament files on the folder "../files/tournament_name"
func SaveTournament(tour_name string, files map[string]string) error {
	for file_name, file_content := range files {
		err := writeFile(tour_name, file_name, file_content)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(tour_name, file_name, file_content string) error {
	file_path := "../files/" + tour_name + "/" + file_name
	//create path
	err := os.MkdirAll(getDir(file_path), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(file_path, []byte(file_content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getDir(file_path string) string {
	return file_path[:strings.LastIndex(file_path, "/")]
}
