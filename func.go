package main

import (
	"fmt"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _:=div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: s%", op)
	}
}

func div(a, b int)(q, r int)  {
	return a / b, a % b
}

func apply(op func(int,int) int, a, b int)int {
	return op(a, b)
}

func pow(a, b int)int{
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ... int) int  {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, error :=eval(13, 2, "!"); error!= nil{
		fmt.Println("Error:", error)
	}else {
		fmt.Println(result)
	}

	fmt.Println(eval(13, 2, "/"))

	q, r :=div(13, 2)

	fmt.Println(q, r)

	fmt.Println(apply(pow, 2, 3))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 3))

	fmt.Println(sum(1,2,3,4,5))
}
