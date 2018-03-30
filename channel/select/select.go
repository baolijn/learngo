package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i ++
		}
	}()
	return out
}

func doWorker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWork(id int) chan<- int {
	c := make(chan int)
	go doWorker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	w := createWork(0)
	n := 0
	hasValue := false
	for  {
		var activeWorker chan<- int
		if hasValue {
			activeWorker = w
		}
		select {
		case n = <- c1:
			hasValue = true
		case n = <- c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}
