package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	time.Sleep(8 * time.Second)
	con, err := net.Dial("tcp", "server:8080")
	if err != nil {
		fmt.Printf("%s %s %s\n", "localhost", "not responding", err.Error())
	} else {
		fmt.Printf("%s %s %s\n", "localhost", "responding on port:", "8080")
	}
	defer con.Close()

	reply := make([]byte, 2048)

	_, err = con.Read(reply)
	if err != nil {
		println("Read from server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server:\n", string(reply))
}
