package main

import (
	"fmt"
	"strings"
)

func main() {
	vowels := []string{"a", "e", "i", "o", "u"}

	var s string
	fmt.Print("Enter string: ")
	fmt.Scan(&s)

	var x string
	for _, char := range s {
		if !contains(vowels, strings.ToLower(string(char))) {
			x += string(char)
		}
	}
	fmt.Println(x)
}

func contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
