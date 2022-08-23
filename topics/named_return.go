package main

import "fmt"

func main() {

	a := 10
	b := 20

	mul, add := mul_div(a, b)

	fmt.Printf("mul is %v, and div is: %v", mul, add)
}

func mul_div(a, b int) (mul, add int) {

	mul = a * b
	add = a + b

	return
}
