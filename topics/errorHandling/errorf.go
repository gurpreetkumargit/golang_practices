package main

import "fmt"

func divide(num1, num2 int) error {

	err := fmt.Errorf("%v can not divide by %v\n", num2, num1)

	if num2 == 0 {
		return (err)
	}
	return nil
}

func main() {
	// age := 100

	// err := fmt.Errorf("age can not be -ve %v\n", age)

	// if age < 0 {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("rohan age is: %v\n", age)
	// }

	val := divide(10, 10)

	if val != nil {
		fmt.Println(val)
	} else {
		fmt.Println("valid")
	}
}
