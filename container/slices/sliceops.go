package slices

import "fmt"

func pringSlice(s []int)  {
	fmt.Printf("%v, len(s)%d, cap(s) %d\n", s, len(s), cap(s))

}
func main() {
	fmt.Println("Creating slice")
	var s []int //Zero value for slice is nil
	for i:=0;i< 100 ;i++  {
		pringSlice(s)
		s = append(s, 2*i +1)
	}
	fmt.Println(s)

	s1 := []int{2, 3, 4, 5}
	pringSlice(s1)
	s2 := make([]int, 16)
	pringSlice(s2)
	s3 := make([]int, 10, 32)
	pringSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	pringSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	pringSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	pringSlice(s2)

	fmt.Println("Popping from tail")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	pringSlice(s2)
}
