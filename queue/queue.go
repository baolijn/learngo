package main

import "fmt"

type queue []int

func (q *queue)push(value int) {
	 *q = append(*q, value)
}

func (q *queue)pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue) isEmpty() bool  {
	return len(*q) == 0
}

func main() {
	qu := queue{1,2}
	fmt.Println(qu)
	qu.push(5)
	qu.push(4)
	fmt.Println(qu)
	qu.pop()
	fmt.Println(qu)
	fmt.Println(qu.isEmpty())
}
