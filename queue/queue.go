package queue

import "fmt"

// An FIFO queue.
type Queue []interface{}

//Pushes the element into the queue.
func (q *Queue)Push(value int) {
	 *q = append(*q, value)
}

// Pops element from head.
func (q *Queue)Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}

func main() {
	qu := Queue{1,2}
	fmt.Println(qu)
	qu.Push(5)
	qu.Push(4)
	fmt.Println(qu)
	qu.Pop()
	fmt.Println(qu)
	fmt.Println(qu.IsEmpty())
}
