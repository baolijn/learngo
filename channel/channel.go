package main

import (
	"fmt"
	"time"
)

func work(id int, c chan int)  {
	func() {
		for{
			fmt.Printf("Worker %d received %c\n", id, <- c)
		}
	}()
}
func chanDemo()  {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go work(i, channels[i])
	}
	for i:= 0; i < 10 ;i++  {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Minute)
}

func main() {
	//c := make(chan int)
	//defer close(c)
	//go func() {c <- 3 + 4}()
	//i := <- c
	//fmt.Println(i)
	chanDemo()
}
