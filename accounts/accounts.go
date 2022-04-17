package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't Withdraw")

// NewAccount : creates Account => struct instance를 만든 후 포인터를 리턴
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit : (a Account)라고 선언해줌으로써 Account struct에 Deposit 메소드가 선언되었다 ===> 메소드를 호출한 인스턴스의 복사본이 아닌 원본에 접근하고 싶은 경우 receiver의 타입을 지정할 때, *을 붙여야 한다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw 입력한 amount 만큼(error는 타입이며 리턴된다)
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner() {}
