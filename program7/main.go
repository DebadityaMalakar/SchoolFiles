package main

import (
	"fmt"
	"strings"
)

func main() {
	var emailID string
	fmt.Print("Enter email id: ")
	fmt.Scan(&emailID)

	parts := strings.Split(emailID, "@")
	if len(parts) != 2 {
		fmt.Println("No data or @ does not exist")
		return
	}

	p1, p2, p3 := parts[0], "@", parts[1]

	if p1 == "" || p2 == "" || p3 == "" {
		fmt.Println("No data or @ does not exist")
	} else if p1 != "" && p2 == "@" && p3 != "edupillar.com" {
		fmt.Println("Domain does not match")
	} else if p1 == "" && p2 == "@" && p3 == "edupillar.com" {
		fmt.Println("Only Domain")
	} else if p1 != "" && p2 == "@" && p3 == "edupillar.com" {
		fmt.Println("Domain Matches")
	} else {
		fmt.Println("Invalid")
	}
}
