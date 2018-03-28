package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

func convertToBin(n int)string {
	result :=""
	for ; n > 0; n /= 2 {
		lsb := n % 2

		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file, err :=os.Open(filename)
	if err != nil {
		panic("")
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	for{
		fmt.Println("sss")
	}
}
func looptest()  {
	for loop := 0; ;{
		fmt.Println(loop)
	}

}
func main() {
	looptest()
	fmt.Println(convertToBin(13),
		convertToBin(0))

	printFile("abc.txt")
	//forever()
	abc := `sdsd
				sdsds
				sdsd
				sd
				
				sdds
				sss`
	printFileContents(strings.NewReader(abc))
}
