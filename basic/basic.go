package basic

import (
	"fmt"
	"math"
)

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

func triangle()  {
	var a, b int = 3, 4

	fmt.Println(calTriangle(a, b))
}

func calTriangle(a, b int) int  {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func consts()  {
	const filename  = "abc.txt"
	const a, b  = 3, 4
	//const a, b  int = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums()  {
	const(
		cpp = iota
		java
		golang
		javascript
	)
	const(
		b = 1 << (10 * iota)
		kb
		tb
	)
	fmt.Println(cpp, java, golang, javascript)
	fmt.Println(b, kb, tb)
}

func main() {
	fmt.Println("Hello go")
	VariableZeroValue()
	VariableInitialVale()
	VariableTypeDeduction()
	VariableShorter()
	fmt.Println(aa, bb, cc)
	triangle()
	consts()
	enums()
}
