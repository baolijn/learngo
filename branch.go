package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}else{
	//	fmt.Printf("%s \n",contents)
	//}
	if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s \n",contents)
	}
}
