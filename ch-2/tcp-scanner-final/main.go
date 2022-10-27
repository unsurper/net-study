package main

import (
	"fmt"
	"net"
	"sort"
)

//工作池
func worker(ports, results chan int) {
	for p := range ports {
		//扫描端口
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//结果为0代表错误
			results <- 0
			continue
		}
		conn.Close()
		//将可以通信的端口压入结果
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	//result管道接收结果
	results := make(chan int)
	var openports []int
	//开启工作池
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	//用协程压入管道进入工作池扫描端口
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		//处理结果,不为零则是开启的端口
		if port != 0 {
			openports = append(openports, port)
		}
	}
	//关闭管道
	close(ports)
	close(results)
	//sort排序
	sort.Ints(openports)
	//打印输出
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
