package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

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
