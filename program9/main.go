package main

import (
	"fmt"
)

func main() {
	var l1 []int

	var n int
	fmt.Print("How many elements do you want? ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Printf("Enter element %d of the list: ", i+1)
		fmt.Scan(&x)
		l1 = append(l1, x)
	}

	maxe := l1[0]
	imax := 0

	for i := 1; i < len(l1); i++ {
		if l1[i] > maxe {
			maxe = l1[i]
			imax = i
		}
	}

	if imax > len(l1)/2 {
		fmt.Println("Maximum element lies in the second half")
	} else if imax < len(l1)/2 {
		fmt.Println("Maximum element lies in the first half")
	}
}
