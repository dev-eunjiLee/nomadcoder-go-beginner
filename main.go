package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/mydict"
)

func main() {

	dictionary := mydict.Dictionary{}

	baseWord := "hello"

	dictionary.Add(baseWord, "First")

	err := dictionary.Update(baseWord, "Second")

	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	deleteWord := "hello"
	dictionary.Delete(deleteWord)
	word2, err2 := dictionary.Search(deleteWord)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("after Delete ", word2)

}
