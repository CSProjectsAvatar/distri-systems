package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"runtime"
)

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

func main() {
	// You can also write it to a file as a whole.
	b := []byte("Hello, World!\n")
	err := os.WriteFile("destination.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// out, err := RunMatch("coin", "p1.py", "p2.py")
	// if err != nil {
	// 	log.Printf("error: %v\n", err)
	// }
	// fmt.Println("--- stdout ---")
	// fmt.Println(out)
}

type GameResult int

const (
	Not_Played GameResult = iota
	P1_Wins
	P2_Wins
	Tie
)
