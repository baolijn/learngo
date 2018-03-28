package main

import "fmt"

func main() {
	x := []int{1,2,3}
	for i,k := range x { x = append(x, k); x = append(x, i)}
	fmt.Println(x)
}
