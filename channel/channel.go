package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int)  {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
}

func createWork(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Minute)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func main() {
	//c := make(chan int)
	//defer close(c)
	//go func() {c <- 3 + 4}()
	//i := <- c
	//fmt.Println(i)
	//chanDemo()
	bufferedChannel()
}
