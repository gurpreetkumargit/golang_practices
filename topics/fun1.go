package main

import "fmt"

func main() {
	fmt.Println("main function")
	greet()
	val := greeter()

	// anonymous func

	func() {
		fmt.Println("anonymous function")
		fmt.Printf("anonymous func creating a closure and accessing data of parent(outer) function: %v\n", val)
	}()

	// assign a anonymous func to a var

	anonymous := func() {
		fmt.Println("2nd anonymous func: assigner using a var")
	}

	anonymous()

}

func greet() {
	fmt.Println("welcome to golang")
}

func greeter() string {
	return "greeter function outside the main() function"
}
