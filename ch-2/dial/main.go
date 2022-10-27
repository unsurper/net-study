package main

import (
	"fmt"
	"net"
)

func main() {
	// net.Dial拨号连接到指定网络上的地址。
	// Examples:
	//	Dial("tcp", "golang.org:http")
	//	Dial("tcp", "192.0.2.1:http")
	//	Dial("tcp", "198.51.100.1:80")
	//	Dial("udp", "[2001:db8::1]:domain")
	//	Dial("udp", "[fe80::1%lo0]:53")
	//	Dial("tcp", ":80")
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
