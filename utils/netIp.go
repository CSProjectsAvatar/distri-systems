package utils

import (
	"log"
	"net"
)

// GetOutboundIP gets preferred outbound ip of this machine.
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetIPString() string {
	return GetOutboundIP().String()
}
