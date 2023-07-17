package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Print("Enter a string: ")
	fmt.Scan(&s)

	if isDigit(s) {
		fmt.Println("The given string is a digit")
	} else if isAlpha(s) {
		fmt.Println("The given string contains alphabets")
	} else {
		fmt.Println("The given string is a special character")
	}
}

func isDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isAlpha(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}
