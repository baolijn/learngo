package main

import "fmt"

func main() {
	m := map[string]string {
		"name" : "eden",
		"course" : "golang",
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

	fmt.Println("Getting value")
	courseName := m["course"]
	fmt.Println(courseName)
	if causeName, ok := m["cause"]; ok{
		fmt.Println(causeName)
	}else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting map")

	delete(m, "site")
	fmt.Println(len(m))

}
