package learn_functions

import "fmt"

// 기본 함수 문법
func Add(a int, b int) int {
	return a + b
}

// 함수의 매개변수 타입을 생략할 수 있음
func Multiply(a, b int) int {
	return a * b
}

// 함수의 반환값이 여러 개일 때
func Divide(a, b int) (int, int) {
	if b == 0 {
		return 0, 0 // 0으로 나누는 경우 처리
	}
	return a / b, a % b // 몫과 나머지 반환
}

// naked function
func Subtract(a, b int) (result int) {
	result = a - b
	return
}

// defer 문법을 사용한 예시
func PrintWithDefer() {
	defer fmt.Println("This is printed last (deferred)")
	fmt.Println("This is printed first")
}
