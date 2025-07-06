package learn_ifelse

import "fmt"

func CheckNumber(n int) {
	if result := n % 2; result == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}
}
