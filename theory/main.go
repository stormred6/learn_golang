package main

import (
	"fmt"
	"learn_golang/theory/learn_functions"
	"learn_golang/theory/learn_ifelse"
	"learn_golang/theory/learn_loop"
	"learn_golang/theory/learn_pointer"
)

type Person struct {
	Name string
	Age  int
}

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

	learn_loop.PrintNumbers(5)                       // Print numbers from 1 to 5
	learn_loop.PrintIndexAndValue([]int{10, 20, 30}) // Print index and value of slice elements
	nums := []int{1, 2, 3, 4, 5}
	sum := learn_loop.SumSlice(nums)  // Sum of slice elements
	fmt.Println("Sum of slice:", sum) // Sum of slice: 15
	learn_loop.WhileLikeLoop(3)       // Demonstrate while-like loop behavior

	learn_ifelse.CheckNumber(10) // 10 is even
	learn_ifelse.CheckNumber(7)  // 7 is odd

	learn_pointer.Swap(&nums[0], &nums[1]) // Swap first two elements
	fmt.Println("After Swap:", nums)       // After Swap: [2 1 3 4 5]

	learn_pointer.SetToZero(&nums[0])     // Set first element to zero
	fmt.Println("After SetToZero:", nums) // After SetToZero: [0 1 3 4 5]

	names := [5]string{"test1", "test2", "test3"}
	fmt.Println("Names:", names) // Names: [test1 test2 test3]

	slice := []string{"apple", "banana", "cherry"}
	fmt.Println("Slice:", slice)              // Slice: [apple banana cherry]
	slice = append(slice, "date")             // Append to slice
	fmt.Println("Slice after append:", slice) // Slice after append: [apple banana cherry date

	nico := map[string]int{"age": 30, "height": 175}
	fmt.Println("Nico:", nico) // Nico: map[age:30 height:175]

	for key, value := range nico {
		fmt.Printf("Key: %s, Value: %d\n", key, value) // Key: age, Value: 30
	}

	youngshin := Person{Name: "Youngshin", Age: 39}
	fmt.Printf("Person: Name=%s, Age=%d\n", youngshin.Name, youngshin.Age) // Person: Name=Youngshin, Age=39
	teayang := Person{"teayang", 39}
	fmt.Printf("Person: Name=%s, Age=%d\n", teayang.Name, teayang.Age) // Person: Name=teayang, Age=39}
}
