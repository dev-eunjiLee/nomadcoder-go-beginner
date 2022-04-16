package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"pizza", "pasta"}
	nico := person{"nico", 18, favFood} // struct 지정하는 방법 1: 순서대로 입력(코드상에서 좋지 못함)
	fmt.Println(nico)
	// struct 지정하는 방법 2: key: value 쌍으로 순서대로 입력(코드에서 명확하게 드러나서 좋다)
	// * 1, 2를 섞어서는 표현 불가
	ethan := person{
		name:    "ethan",
		age:     28,
		favFood: favFood,
	}

	v := reflect.ValueOf(ethan)
	test := v.Type()

	for i := 0; i < test.NumField(); i++ {
		fmt.Println(i, test.Field(i).Name, v.Field(i))
	}

	fmt.Println(v, test)
}
