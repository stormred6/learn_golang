package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errorInsufficientAccount = errors.New("insufficient account balance")

// NewAccount 함수는 새로운 은행 계좌를 생성합니다.
func NewAccount(owner string) *Account {
	return &Account{
		owner:   owner,
		balance: 0,
	}
}

// Balance 함수는 계좌의 현재 잔액을 반환합니다.
func (a *Account) Balance() int {
	return a.balance
}

// Deposit 함수는 계좌에 금액을 입금합니다.
func (a *Account) Deposit(amount int) {
	if amount < 0 {
		return // 음수 금액은 입금하지 않습니다.
	}
	a.balance += amount
}

// Withdraw 함수는 계좌에서 금액을 출금합니다.
func (a *Account) Withdraw(amount int) error {
	if amount < 0 || amount > a.balance {
		return errorInsufficientAccount
	}
	a.balance -= amount

	return nil
}

// Owner 함수는 계좌의 소유자를 반환합니다.
func (a *Account) Owner() string {
	return a.owner
}

// String 함수는 계좌 정보를 문자열로 반환합니다.
func (a *Account) String() string {
	return a.Owner() + "'s account" + " has a balance of " + fmt.Sprintf("%d", a.Balance())
}
