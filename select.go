package main

import (
	"fmt"
	"runtime"
	"time"
)

var ch1 = make(chan int)

// 对nil管道执行操作将会一直阻塞！
var ch2 chan int
var chs = []chan int{ch1, ch2}
var nums = []int{1, 2, 3, 4, 5}

func main() {
	runtime.GOMAXPROCS(4)
	go func(ch chan int) {
		fmt.Println("go routine start")
		<-ch
	}(ch1)
	time.Sleep(1 * time.Second)
	select {
	// 无缓冲的通道，只有发送和接收同时准备好才能进行通信
	case ch1 <- 1:
		fmt.Println("case 0 is selected")
	default:
		fmt.Println("defalut is selected")
	}
}

func getChan(i int) chan int {
	fmt.Println("getChan", i)
	return chs[i]
}

func getNumber(i int) int {
	fmt.Println("getNumber", i)
	return nums[i]
}
