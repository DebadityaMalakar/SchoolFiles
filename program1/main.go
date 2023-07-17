package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Print("Enter number: ")
	_, err := fmt.Scanf("%f", &n)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if (n > 0 && n < 10) || (n < 0 && n > -10) {
		fmt.Printf("%.2f is a 1-digit number\n", n)
	} else if (n >= 10 && n <= 99) || (n <= -10 && n >= -99) {
		fmt.Printf("%.2f is a 2-digit number\n", n)
	} else if (n > 99 && n <= 999) || (n < -99 && n >= -999) {
		fmt.Printf("%.2f is a 3-digit number\n", n)
	} else {
		fmt.Println("Invalid Case")
	}
}
