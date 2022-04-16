package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 200
	fmt.Println(a)
}
