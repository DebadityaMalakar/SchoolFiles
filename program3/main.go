package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if num > 1 {
		i := 2
		for i <= int(math.Sqrt(float64(num))) {
			if num%i == 0 {
				fmt.Printf("%d is not a prime number\n", num)
				break
			}
			i++
		}

		if i > int(math.Sqrt(float64(num))) {
			fmt.Printf("%d is a prime number\n", num)
		}
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}
}
