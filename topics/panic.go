package main

import "fmt"

func divide(num1, num2 int) int {

	if num2 == 0 {
		panic("value can not be zero")
	} else {
		var op = (num1 / num2)
		return (op)
	}

}

func main() {

	// panic statement immediately ends the execution of the programs

	// fmt.Println("start")
	// fmt.Println("processing")
	// // panic("stop")
	// fmt.Println("end")

	fmt.Println(divide(10, 3))
	fmt.Println(divide(2, 0))

}
