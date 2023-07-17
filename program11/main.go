package main

import (
	"fmt"
	"strings"
)

func main() {
	var L []string
	var u []string
	var dname []string

	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x string
		fmt.Printf("Enter email id of student %d: ", i+1)
		fmt.Scan(&x)

		for !strings.Contains(x, "@") {
			fmt.Println("Invalid email id")
			fmt.Print("Please enter a valid email id: ")
			fmt.Scan(&x)
		}

		L = append(L, x)
	}

	for _, j := range L {
		U := strings.Split(j, "@")[0]
		Dn := strings.Split(j, "@")[1]
		u = append(u, U)
		dname = append(dname, Dn)
	}

	fmt.Println(u)
	fmt.Println(dname)
}
