package main

import (
	"fmt"
)

type s struct {
	a string
	b string
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	sum := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)
}
