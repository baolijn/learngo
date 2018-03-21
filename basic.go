package main

import "fmt"

var (
	aa = 3
	bb = "qqq"
	cc = true
)
func VariableZeroValue()  {
	var a int
	var abc string
	fmt.Printf("%d %q",a, abc)
}

func VariableInitialVale()  {
	var a, c int = 1,3
	var abc string = "abc"
	fmt.Println(a, abc, c)
}

func VariableTypeDeduction()  {
	var a, b, c, d = 1,3, true, "abc"
	fmt.Println(a, b, c, d)
}

func VariableShorter()  {
	a, b, c, d :=1,3,true,"def"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("Hello go")
	VariableZeroValue()
	VariableInitialVale()
	VariableTypeDeduction()
	VariableShorter()
	fmt.Println(aa, bb, cc)
}
