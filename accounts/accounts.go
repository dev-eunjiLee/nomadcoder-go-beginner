package accounts

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
