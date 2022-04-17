package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("nico") // account struct를 nico를 이용해 instance로 만든 후, 해당 인스턴스의 주소값을 리턴
	account.Deposit(3)
	fmt.Println(account.Balance())
	err := account.Withdraw(5)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err) // log.FatalLn(): 프로그램을 종료시킨다
	}
	fmt.Println(account.Balance())
}
