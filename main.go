package main

import (
	"fmt"
	"study/learn_golang/learn_functions"
)

func main() {
	fmt.Println("Add:", learn_functions.Add(3, 5))           // Add: 8
	fmt.Println("Multiply:", learn_functions.Multiply(4, 6)) // Multiply: 24

	// divide 함수 호출
	quotient, remainder := learn_functions.Divide(10, 3)
	fmt.Println("Divide:", quotient, "Remainder:", remainder) // Divide: 3 Remainder: 1

	// 0으로 나누는 경우
	q, r := learn_functions.Divide(10, 0)
	fmt.Println("Divide by zero:", q, "Remainder:", r) // Divide by zero: 0 Remainder: 0

	// 나머지값은 무시하고자 할 때
	q, _ = learn_functions.Divide(10, 3)
	fmt.Println("Divide ignoring remainder:", q) // Divide ignoring remainder: 3

	fmt.Println("Subtract with naked return:", learn_functions.Subtract(10, 4)) // Subtract with naked return: 6

	learn_functions.PrintWithDefer() // This will demonstrate the defer behavior
}
