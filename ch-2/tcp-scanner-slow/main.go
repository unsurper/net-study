package main

import (
	"fmt"
	"net"
)

func main() {
	//端口扫描1-1024
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		//等待响应耗时长
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d closed\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
