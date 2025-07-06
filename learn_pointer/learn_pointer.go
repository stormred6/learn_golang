package learn_pointer

import "fmt"

// Swap 함수는 두 개의 int 포인터 a와 b가 가리키는 값들을 서로 교환합니다.
// 이 함수는 포인터를 사용하므로, 함수 호출 이후에도 원본 변수의 값이 실제로 바뀌게 됩니다.
// 예를 들어, 두 변수의 주소를 인자로 넘기면 두 변수의 값이 서로 바뀝니다.
//
// 예시:
//
//	x := 10
//	y := 20
//	Swap(&x, &y)
//	// x는 20, y는 10이 됩니다.
func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func SetToZero(ptr *int) {
	fmt.Printf("Address of ptr: %p, Value: %d\n", ptr, *ptr)
	*ptr = 0
}
