package main

import "fmt"

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("RECOVER()", a)
	}
}

func divide(num1, num2 int) int {

	defer handlePanic()

	if num2 == 0 {
		panic("value can not be zero")
	} else {
		op := num1 / num2
		return (op)
	}
}

func main() {
	// recover statement
	// the panic statement immediately terminate the execution of programs but in some case its important for a program to be executed.const
	// so panic statement recovers the programs after the termination of panic

	fmt.Println("op", divide(10, 4))
	fmt.Println("op", divide(2, 0))
}
