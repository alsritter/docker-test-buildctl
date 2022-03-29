package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	fmt.Println("IPv4: ", GetOutboundIP())

	fmt.Println("开始监听 9999 端口")
	http.ListenAndServe(":9999", nil)
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
