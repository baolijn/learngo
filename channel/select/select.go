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
	var values []int
	var activeValue int
	for  {
		var activeWorker chan<- int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <- c1:
			values = append(values, n)
		case n = <- c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		}
	}
}
