package main

import "fmt"

func main() {
	// var slice = []int{1, 3, 4, 5}

	// fmt.Print(slice)

	newSlice := []int{1, 2, 3, 6}

	fmt.Println(newSlice)

	newSlice = append(newSlice, 7)

	fmt.Println(newSlice)

	array := [5]int{1, 2, 4, 6, 7}

	// mySlice := array[start:end]

	mySlice := array[2:5]
	fmt.Print("created slice from array: ")
	fmt.Println(mySlice)
}
