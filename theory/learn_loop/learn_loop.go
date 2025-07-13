package learn_loop

import "fmt"

// PrintNumbers prints numbers from 1 to n using a for loop.
func PrintNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

// SumSlice sums all elements in a slice using a for loop.
func SumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// PrintIndexAndValue prints both index and value of elements in a slice using a range loop.
func PrintIndexAndValue(nums []int) {
	for idx, val := range nums {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
	}
}

// WhileLikeLoop demonstrates a while-like loop in Go.
func WhileLikeLoop(limit int) {
	i := 0
	for i < limit {
		fmt.Println("i is", i)
		i++
	}
}
