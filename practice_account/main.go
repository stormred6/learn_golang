package main

import (
	"fmt"
	"learn_golang/practice_account/accounts"
)

func main() {
	accounts := accounts.NewAccount("youngshin")
	fmt.Println("Account Owner:", accounts) // Account Owner: John Doe

	accounts.Deposit(1000)
	fmt.Println("Balance after deposit:", accounts.Balance()) // Balance after deposit: 1000

	err := accounts.Withdraw(1500)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	} else {
		fmt.Println("Balance after withdrawal:", accounts.Balance()) // Balance after withdrawal: 500
	}
}
