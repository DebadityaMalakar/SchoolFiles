package main

import (
	"fmt"
)

func convertCurrency(amtInDls, conRate float64) float64 {
	conAmt := amtInDls * conRate
	fmt.Println(conAmt)
	return conAmt
}

func main() {
	var amt float64
	fmt.Print("Enter amount in $: ")
	fmt.Scan(&amt)

	var cr float64
	fmt.Print("Enter conversion rate: ")
	fmt.Scan(&cr)

	cnvrted := convertCurrency(amt, cr)
	fmt.Printf("₹ %.2f\n", cnvrted)
}
