package main

import (
	"fmt"
	"net"
	"os"

	"github.com/CSProjectsAvatar/distri-systems/utils"
)

func main() {
	con, err := net.Dial("tcp", ":8000")
	utils.CheckErr(err)
	defer con.Close()

	_, err = con.Write([]byte(os.Args[1] + "\n")) // sends the student's name to server
	utils.CheckErr(err)

	reply := make([]byte, 1024)
	_, err = con.Read(reply) // listen for server answer
	utils.CheckErr(err)

	fmt.Print(string(reply)) // printing answer
}
