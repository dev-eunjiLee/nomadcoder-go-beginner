package main

import "fmt"

func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	// * for문을 이용해 map 열거
	for key, value := range nico {
		fmt.Println(key, value)
	}
	fmt.Println(nico)
}
