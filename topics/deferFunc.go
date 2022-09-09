package main

// the defer delays the execution of func, method, or anony func until the nearby methods return

import "fmt"

func mul(a, b int) {
	fmt.Printf("result: %v\n", a*b)
}

func show() {
	fmt.Println("we are learning defer")
}

// we can create multiple defers in our programs
// defers work in LIFO order (in 2nd code example)
// generally used to ensure that files are closed when their need is over, or to close the channel

// func main() { // example 1

// 	mul(10, 10)

// 	defer mul(20, 20)

// 	mul(30, 30)

// 	show()

// }

func main() { // example 2

	fmt.Println("start") // 1st op

	defer fmt.Println("end") // end  (LIFO)
	defer mul(10, 10)
	defer fmt.Println("2nd execution: ")

	defer mul(20, 20)
	defer fmt.Println("1st execution: ") // execution will be start here

}
