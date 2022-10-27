package main

import (
	"fmt"
	"sync"
)

//work工作池
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	//管道最多堵塞100个协程
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	//cap(ports),ports管道的容量开启线程
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		//将i压入管道内随机分配到工作池中
		ports <- i
	}
	wg.Wait()
	close(ports)
}
