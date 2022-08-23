package main

import "fmt"

func main() {

	// var name string
	// fmt.Println("enter your name: ")

	// fmt.Scan(&name)

	// var age int
	// fmt.Println("enter you age: ")
	// fmt.Scan(&age)

	// fmt.Printf("user %v age is %v ", name, age)

	var (
		name    string
		age     int
		married bool
	)
	ag1, err := fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married) //  The identification format must be filled in before the parameters in this function, such as 1: 2: ...
	fmt.Println(ag1)
	fmt.Println(err)
	fmt.Printf(" Scan results  name:%s age:%d married:%t", name, age, married)
}
