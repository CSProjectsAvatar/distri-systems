package main

import (
	"fmt"
	"net"

	"github.com/CSProjectsAvatar/distri-systems/utils"
)

func main() {
	con, err := net.Dial("tcp", "www.uh.cu:80")
	utils.CheckErr(err)
	defer con.Close()

	_, err = con.Write([]byte("GET / HTTP/1.1\r\nHost: www.uh.cu\r\n\r\n"))
	utils.CheckErr(err)

	reply := make([]byte, 1024)
	_, err = con.Read(reply)
	utils.CheckErr(err)

	fmt.Println(string(reply))
}
