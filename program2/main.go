package main

import "fmt"

func main() {
	for i := 10; i <= 15; i++ {
		fmt.Printf("Cube of %d - %d\n", i, i*i*i)
	}
}
