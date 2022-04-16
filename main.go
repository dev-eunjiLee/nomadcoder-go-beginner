package main

import "fmt"

func main() {
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "3"
	names[4] = "4"

	slice := []string{"nico", "lynn", "dal"}
	newSlice := append(slice, "3")

	fmt.Println(names)
	fmt.Println(slice)
	fmt.Println(newSlice)

}
