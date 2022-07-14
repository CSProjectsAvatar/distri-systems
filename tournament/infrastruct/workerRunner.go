package infrastruct

import (
	"bytes"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
)

type WorkerRunner struct {
	dm usecases.DataMngr
}

func NewWorkerRunner(dm usecases.DataMngr) *WorkerRunner {
	return &WorkerRunner{dm}
}
func (wr *WorkerRunner) RunMatch(match *domain.Pairing) (domain.MatchResult, error) {
	tInfo, err := wr.dm.GetTournInfo(match.TourId) // get tournament info
	if err != nil {
		return domain.NotPlayed, err
	}
	// calling to FileGroup
	fileNames := []string{tInfo.Name, match.Player1.Id, match.Player2.Id}
	files, err := wr.dm.FileGroup(match.TourId, fileNames)
	if err != nil {
		return domain.NotPlayed, err
	}

	folderName := tInfo.Name + tInfo.ID
	SaveFilesLocal(folderName, files)
	// Run the match
	//return RunMatch(folderName, tInfo.Name, match.Player1.Id, match.Player2.Id)
	// rand in [1,3]
	r := rand.Intn(3) + 1
	return domain.MatchResult(r), nil
}

func SaveFilesLocal(folderName string, files map[string]string) error {
	for fileName, fileContent := range files {
		err := writeFile(folderName, fileName, fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(folderName, file_name, file_content string) error {
	file_path := "../files/" + folderName + "/" + file_name
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

// Run a match between two players by calling the game with them as parameters
func RunMatch(folderName, tName, player_i, player_j string) (domain.MatchResult, error) {
	command := "python ../files/" + folderName + "/" + tName + ".py " + player_i + ".py " + player_j + ".py"
	err, out, _ := RunCommand(command)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	winner, err := strconv.Atoi(out[0:1]) // for deal with "2\r\n" response styles
	if err != nil {

		log.Printf("Couldn parse response (%v), from %v\n error: %v", winner, tName, err)
	}
	return domain.MatchResult(winner), nil
}

func getDir(file_path string) string {
	return file_path[:strings.LastIndex(file_path, "/")]
}

// Run a specific comand in the terminal (cmd or bash)
func RunCommand(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	os := runtime.GOOS
	var shellToUse, fArg string
	switch os {
	case "windows":
		shellToUse = "cmd"
		fArg = "/C"
	default:
		shellToUse = "bash"
		fArg = "-c"
	}
	cmd := exec.Command(shellToUse, fArg, command)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
