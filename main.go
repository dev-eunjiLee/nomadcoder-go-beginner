package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(&dictionary)
	definition, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
