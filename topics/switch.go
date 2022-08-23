package main

import "fmt"

func main() {

	// expression switch
	var day int = 2

	switch day {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

}
