package main

import (
	"fmt"
	"os"
	"bufio"
	"learngo/functional/fib"
)

func tryDefer()  {
	for i :=0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string)  {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil{
		if pathError, ok := err.(*os.PathError); !ok{
			panic("File does not exist")
		}else {
			fmt.Println(pathError.Op,
				pathError.Path,
					pathError.Err)
		}

	}
	defer file.Close()

	writer :=bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0;i < 20 ;i ++  {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
