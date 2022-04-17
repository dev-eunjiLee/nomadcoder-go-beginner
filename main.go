package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/mydict"
)

func main() {

	word := "first"
	def := "First word"

	dictionary := mydict.Dictionary{word: def}

	word = "name"
	def = "ethan"

	addErr := dictionary.Add(word, def)

	if addErr != nil {
		fmt.Println(addErr)
	}

	definition, err := dictionary.Search("name")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
