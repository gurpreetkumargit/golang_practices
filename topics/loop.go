package main

import "fmt"

func main() {

	// slice := []int{1, 2, 3, 4, 5}

	// index := 0
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Print(slice[index], " ")
	// 	index++
	// }

	// for loop as infinite loop

	// for {
	// 	fmt.Println("loop")
	// }

	// for loop as while loop

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// range in for loop

	// for i, j := range slice {
	// 	fmt.Println(i, j)
	// }

	//use loop for String

	str := "abcd"

	for index, char := range str {
		fmt.Println(string(char), index)
	}

}
