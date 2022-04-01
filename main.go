package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("Hello World!"))
}

func main() {
	// 打印 IP 地址
	fmt.Println("IPv4: ", GetOutboundIP())

	// 本地请求外网
	resp, _ := http.Get("http://baidu.com")
	fmt.Printf("%v \n", resp.Status)

	// 本地起端口不冲突
	var s server
	fmt.Println("开始监听 9999 端口")
	http.ListenAndServe(":9999", &s)
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
