package main

import "fmt"

func main() {

	// if statement
	a := 10

	// if else
	if a == 10 {
		fmt.Println("string")
	} else {
		fmt.Println("not a string")
	}

	b := 100

	// nested is cond.
	if b == 100 {
		if a == 10 {
			fmt.Println("nested if")
		}
	}

	// if else ladder
	if b == a {
		fmt.Println("equal")
	} else if b > a {
		fmt.Println("greater")
	} else {
		fmt.Println("less")
	}
}
