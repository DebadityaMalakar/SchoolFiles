package main

import "fmt"

func main() {
	var s string
	fmt.Print("Enter a string: ")
	fmt.Scan(&s)

	if isPalindrome(s) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}

func isPalindrome(s string) bool {
	r := []rune(s)
	length := len(r)

	for i := 0; i < length/2; i++ {
		if r[i] != r[length-i-1] {
			return false
		}
	}

	return true
}
