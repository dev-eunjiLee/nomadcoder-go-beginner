package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico") // account struct를 nico를 이용해 instance로 만든 후, 해당 인스턴스의 주소값을 리턴
	fmt.Println(account)
}
