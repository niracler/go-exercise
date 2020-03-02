package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(10, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}

func chanDemo() {
	var channels [10]chan<- int

	// 要有一个 goroutine 接收数据
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(11, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	bufferedChannel()
	channelClose()
}
