package main

import "fmt"

func updateSlice(arr []int)  {
	arr[0] = 100
}
func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}

	fmt.Println("arr[2:6] is:", arr[2:6])
	fmt.Println("arr[:6] is:", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 is:", s1)
	s2 := arr[:]
	fmt.Println("s2 is:", s2)
	updateSlice(s1)
	fmt.Println("After updateSlice:")
	fmt.Println("s1 is:", s1)
	updateSlice(s2)
	fmt.Println("s2 is:", s2)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending Slice")
	arr[0] = 0
	arr[2] = 2
	fmt.Println(arr)

	s1 = arr[2:6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 = s1[3:5]
	fmt.Println(s2)
}
