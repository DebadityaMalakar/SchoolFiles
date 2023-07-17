package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i == 3 {
			i++
			continue
		}
		fmt.Println(i)
		i++
	}
}
