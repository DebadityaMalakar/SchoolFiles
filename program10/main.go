package main

import (
	"fmt"
)

func main() {
	var L []int

	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Printf("Enter element %d of the tuple: ", i+1)
		fmt.Scan(&x)
		L = append(L, x)
	}

	L11 := make([]int, 0)
	for i := len(L) - 1; i > 0; i -= 3 {
		L11 = append(L11, L[i])
	}

	L12 := make([]int, 0)
	for i := 2; i < 10 && i < len(L); i += 2 {
		L12 = append(L12, L[i])
	}

	fmt.Println(L11)
	fmt.Println(L12)
}
