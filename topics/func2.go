package main

import (
	"fmt"
	"strings"
)

// variadic function

func main() {

	add := adder(2, 4)

	fmt.Printf("adder func result: %v\n", add)

	// variadic function using int value
	proResult := proAdder(1, 3, 4, 87, 5, 7, 5)

	fmt.Printf("proAdder func result: %v\n", proResult)

	// variadic function using string value

	resProString := proString("p", "r", "e", "e", "t")

	fmt.Printf("result of proString %v\n", resProString)

	// pass a slice in variadic function

	nameString := []string{"Preet", "Dhiman", "Kumar"}

	fmt.Println("result by passing slice in variadic function: ")
	fmt.Printf(proString(nameString...))
}

func adder(one int, two int) int {
	return one + two
}

// giving infinite no. of parameters in our function
func proAdder(values ...int) int {
	total := 0

	// all values will be come as a slice so we will use range loop to add all value in total var

	for _, val := range values {
		total += val
	}

	return total
}

func proString(elements ...string) string {
	return (strings.Join(elements, "-"))
}
