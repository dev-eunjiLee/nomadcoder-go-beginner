package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/something"
)

func main() {
	fmt.Println("hello world")

	something.SayHello()
	//something.sayBye() => 소문자로 시작하는 함수는 외부에서 호출 불가
}
