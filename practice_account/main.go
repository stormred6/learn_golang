package main

import (
	"fmt"
	"learn_golang/practice_account/accounts"
	"learn_golang/practice_account/dictionary"
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

	accounts.ChangeOwner("Jane Doe")

	dictionary := dictionary.Dictionary{}

	dictionary["hello"] = "world"
	fmt.Println("Dictionary entry:", dictionary["hello"]) // Dictionary entry: world

	value, err := dictionary.Search("test")
	if err != nil {
		fmt.Println("Search error:", err)
	} else {
		fmt.Println("Search result:", value) // Search result: world
	}

	err = dictionary.Add("test", "new world")
	if err != nil {
		fmt.Println("Add error:", err) // Add error: word already exists
	} else {
		fmt.Println("Added new entry to dictionary")
	}

	err = dictionary.Update("hello", "everyone")
	if err != nil {
		fmt.Println("Update error:", err) // Update error: word not found
	} else {
		fmt.Println("Updated dictionary entry")
		fmt.Println(dictionary["hello"]) // Updated dictionary entry: everyone
	}

	dictionary.Delete("hello")
	value, err = dictionary.Search("hello")
	if err != nil {
		fmt.Println("Search after delete error:", err) // Search after delete error: word not found
	} else {
		fmt.Println("Search result after delete:", value) // Search result after delete:
	}
}
