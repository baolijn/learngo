package main

import "fmt"

func main() {
	m := map[string]string {
		"name" : "eden",
		"cource" : "golang",
		"site" : "golang.org",
		"quality" : "not bad",
	}

	m2 := make(map[string]string)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("Travering map")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
