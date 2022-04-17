package accounts

import "fmt"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount : creates Account => struct instance를 만든 후 포인터를 리턴
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit : (a Account)라고 선언해줌으로써 Account struct에 Deposit 메소드가 선언되었다
func (a Account) Deposit(amount int) {
	fmt.Println(amount)
}
