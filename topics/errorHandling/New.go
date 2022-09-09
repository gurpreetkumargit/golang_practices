package main

import (
	"errors"
	"fmt"
)

func checkName(name string) error {
	newErr := errors.New("error came")

	if name != "preet" {
		return newErr
	}

	return nil
}

func main() {

	// message := "heo"

	// myErr := errors.New("wrong message")

	// if message != "hello" {
	// 	fmt.Println(myErr)
	// } else {
	// 	fmt.Println(message)
	// }

	name := "preet"

	err := checkName(name)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("valid name")
	}

	// fmt.Println(err)
}
