package main

import "fmt"

// 기본 함수 문법
func add(a int, b int) int {
	return a + b
}

// 함수의 매개변수 타입을 생략할 수 있음
func multiply(a, b int) int {
	return a * b
}

// 함수의 반환값이 여러 개일 때
func divide(a, b int) (int, int) {
	if b == 0 {
		return 0, 0 // 0으로 나누는 경우 처리
	}
	return a / b, a % b // 몫과 나머지 반환
}

func main() {
	fmt.Println("Add:", add(3, 5))           // Add: 8
	fmt.Println("Multiply:", multiply(4, 6)) // Multiply: 24

	// divide 함수 호출
	quotient, remainder := divide(10, 3)
	fmt.Println("Divide:", quotient, "Remainder:", remainder) // Divide: 3 Remainder: 1

	// 0으로 나누는 경우
	q, r := divide(10, 0)
	fmt.Println("Divide by zero:", q, "Remainder:", r) // Divide by zero: 0 Remainder: 0

	// 나머지값은 무시하고자 할 때
	q, _ = divide(10, 3)
	fmt.Println("Divide ignoring remainder:", q) // Divide ignoring remainder: 3
}
