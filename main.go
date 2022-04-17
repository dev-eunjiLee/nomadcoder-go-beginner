package main

import (
	"fmt"
	"github.com/dev-eunjiLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico") // account struct를 nico를 이용해 instance로 만든 후, 해당 인스턴스의 주소값을 리턴

	// 포인터를 역참조해서 필드에 접근할 수는 없지만, 역참조한 인스턴스를 다른 변수에 할당해 해당 인스턴스를 수정하는 방식은 가능하다.
	newAccount := *account
	newAccount.Owner = "nico2"

	fmt.Println(&newAccount, account) // newAccount
	fmt.Println(newAccount)           // {nico2 0}

	test := accounts.Account{Owner: "test"}
	test.Owner = "test2"
	fmt.Println(test)

	fmt.Println(account)
}
