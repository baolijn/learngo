package fib

//F(0)=0，F(1)=1, F(n)=F(n-1)+F(n-2)（n>=2，n∈N*）
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a +b
		return a
	}
}
