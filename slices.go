package main

import "fmt"

func main() {
	arr := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println("arr[2:6] is:", arr[2:6])
	fmt.Println("arr[:6] is:", arr[:6])
	fmt.Println("arr[2:] is:", arr[2:])
	fmt.Println("arr[:] is:", arr[:])
}
